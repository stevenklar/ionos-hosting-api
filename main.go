package main

import (
	"context"
	"fmt"

	"github.com/stevenklar/ionos-hosting-api/pkg"
)

func main() {
	auth := context.WithValue(context.Background(), pkg.ContextAPIKey, pkg.APIKey{
		Key:    "",
		Prefix: "",
	})

	cfg := pkg.NewConfigurationDns()
	client := pkg.NewAPIClient(cfg)
	zones, _, err := client.ZonesApi.GetZones(auth)
	if err != nil {
		fmt.Printf("[ZonesApi] GetZones: %s", err)
	}

	for _, zone := range zones {
		fmt.Printf("%s = %s\n", zone.Id, zone.Name)
	}
}
