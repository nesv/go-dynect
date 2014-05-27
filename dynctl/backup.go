/*
Copyright 2014 Nick Saika

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

// This file holds the implementation of the "backup" sub-command.

import (
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/nesv/go-dynect/dynect"
)

var BackupDescription = `
   Backup the specified zones (or all zones if none are specified), and write
   the contents out as BIND-friendly zone files.`

// Function Backup implements the "dynctl backup" command.
//
// The "dynctl backup" command retrieves all zones managed by DynECT, or a
// specified subset as extra arguments to the "backup" command, and writes out
// BIND-friendly zone files in the specified directory.
func Backup(c *cli.Context) {
	// Make sure an output directory was specified.
	outputDir := c.String("outputdir")
	if outputDir == "" {
		log.Fatalln("no output directory specified")
	}

	// Gather the zones.
	zones, err := listZones()
	if err != nil {
		log.Fatalln(err)
	}

	if c.GlobalBool("verbose") {
		log.Printf("available zones: %v", zones)
	}

	// See if the user only wants to back up one zone, or all zones.
	var zonesToDump = make([]string, 0)
	if len(c.Args()) == 0 {
		// All zones.
		for _, z := range zones {
			parts := strings.Split(z, "/")
			zname := parts[len(parts)-2]
			zonesToDump = append(zonesToDump, zname)
		}
	} else {
		// Only the zones the user specified.
		for _, uz := range c.Args() {
			// Loop over the zones the API says we have, and try
			// to match the zone the user specified with those
			// Dyn says we have.
			for _, z := range zones {
				parts := strings.Split(z, "/")
				zoneName := parts[len(parts)-2]
				if zoneName == uz {
					zonesToDump = append(zonesToDump, zoneName)
					if c.GlobalBool("verbose") {
						log.Printf("queuing zone %q", z)
					}
				}
			}
		}
	}

	// Make sure there is, at least, one zone to backup.
	if len(zonesToDump) == 0 {
		log.Fatalln("no zones to dump")
	}

	zoneData := make(map[string]dynect.ZoneDataBlock, 0)
	for _, zone := range zonesToDump {
		uri := strings.Join([]string{"Zone", zone}, "/")
		var zoneresp dynect.ZoneResponse
		if err := Dyn.Do("GET", uri, nil, &zoneresp); err != nil {
			log.Fatalln(err)
		}

		zoneData[zone] = zoneresp.Data

		if c.GlobalBool("verbose") {
			log.Printf("%s\n\t%#v", zone, zoneresp.Data)
		}
	}

	// Fetch all records in the zone.
	var zoneRecords = make(map[string]dynect.AllRecordsResponse)
	for _, zone := range zonesToDump {
		uri := strings.Join([]string{"AllRecord", zone}, "/")
		var resp dynect.AllRecordsResponse
		if err := Dyn.Do("GET", uri, nil, &resp); err != nil {
			log.Fatalln(err)
		}

		if c.GlobalBool("verbose") {
			log.Printf("%d records in zone %s", len(resp.Data), zone)
			log.Printf("\t%#v", resp.Data)
		}

		zoneRecords[zone] = resp
	}

	var records = make(map[string][]dynect.RecordResponse)
	for zone, zrecs := range zoneRecords {
		zrs := make([]dynect.RecordResponse, 0)

		for _, rec := range zrecs.Data {
			parts := strings.Split(rec, "/")

			fqdn := parts[len(parts)-2]
			id := parts[len(parts)-1]

			uri := strings.Join([]string{parts[2], zone, fqdn, id}, "/")

			var resp dynect.RecordResponse
			if err := Dyn.Do("GET", uri, nil, &resp); err != nil {
				log.Fatalln(err)
			}

			zrs = append(zrs, resp)

			if c.GlobalBool("verbose") {
				log.Printf("retrieved record data for %s/%s/%s", zone, fqdn, id)
			}
		}

		records[zone] = zrs
	}
	return
}
