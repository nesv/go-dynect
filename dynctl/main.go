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
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/fhs/go-netrc/netrc"
	"github.com/nesv/go-dynect/dynect"
)

var (
	Customer string
	Username string
	Password string
	Dyn      *dynect.Client
)

const (
	VERSION = "0.1.0"
)

func main() {
	log.SetFlags(0)
	app := cli.NewApp()
	app.Name = "dynctl"
	app.Usage = "utility for working with DynECT's API"
	app.Version = VERSION
	app.Before = Initialize
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "customer, c", Value: "", Usage: "DynECT customer name"},
		cli.StringFlag{Name: "machine, m", Value: "api.dynect.net", Usage: "netrc machine name to refer to for credentials"},
		cli.StringFlag{Name: "zone, z", Value: "", Usage: "the DNS zone"},
		cli.BoolFlag{Name: "porcelain, p", Usage: "enable porcelain output for piping"}}
	app.Commands = []cli.Command{
		{Name: "list-zones", ShortName: "lsz", Usage: "list all zones", Action: ListZones}}
	app.Run(os.Args)
}

func Initialize(c *cli.Context) error {
	if c.String("customer") == "" {
		err := fmt.Errorf("you must specify a customer name")
		log.Println(err)
		return err
	} else {
		Customer = c.String("customer")
	}

	// Load credentials from ~/.netrc.
	pth := filepath.Join(os.Getenv("HOME"), ".netrc")
	if machines, _, err := netrc.ParseFile(pth); err != nil {
		return err
	} else if len(machines) == 0 {
		return fmt.Errorf("no machines defined in %v", pth)
	} else {
		for _, machine := range machines {
			if machine.Name == c.String("machine") {
				Dyn = dynect.NewClient(Customer)
				err := Dyn.Login(machine.Login, machine.Password)
				if err != nil {
					return err
				}
				return nil
			}
		}
		return fmt.Errorf("no machine with name %v", c.String("machine"))
	}
}
