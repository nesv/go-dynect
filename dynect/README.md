# go-dynect/dynect

Package `github.com/nesv/go-dynect/dynect` provides an HTTP/JSON client and
struct for working with DynECT's API.

## Example usage

The following example retrieves a list of all zones in your managed DNS.

	package main

	import (
			"log"
			"strings"
			
			"github.com/nesv/go-dynect/dynect"
	)

	func main() {
		client := dynect.NewClient("my-dyn-customer-name")
		err := client.Login("my-dyn-username", "my-dyn-password")
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			err := client.Logout()
			if err != nil {
				log.Fatal(err)
			}
		}()

		// Make a request to the API, to get a list of all, managed DNS zones
		var response ZonesResponse
		if err := client.Do("GET", "Zone", nil, &response); err != nil {
			log.Println(err)
		}

		for _, zoneURI := range response.Data {
			zparts := strings.Split(zoneURI, "/")
			zoneName := zparts[len(zparts)-2]
			log.Println(zoneName)
		}
	}

More to come!
