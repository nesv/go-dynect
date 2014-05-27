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

import (
	"fmt"
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/nesv/go-dynect/dynect"
)

func listZones() (zones []string, err error) {
	var zonesResponse dynect.ZonesResponse
	err = Dyn.Do("GET", "Zone", nil, &zonesResponse)
	zones = zonesResponse.Data
	return
}

func ListZones(c *cli.Context) {
	zones, err := listZones()
	if err != nil {
		log.Fatalln(err)
	}

	for _, zone := range zones {
		parts := strings.Split(zone, "/")
		plen := len(parts)
		fmt.Println(parts[plen-2])
	}
}
