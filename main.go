package main

import (
	"context"
	"fmt"

	ionos "github.com/stevenklar/ionos-hosting-api/pkg"
)

func main() {
	auth := context.WithValue(context.Background(), ionos.ContextAPIKey, ionos.APIKey{
		Key:    "",
		Prefix: "",
	})

	dnsExample(auth)
	sslExample(auth)
}

func sslExample(auth context.Context) {
	fmt.Println("========= SSL =========")
	ssl := ionos.NewAPIClient(ionos.NewConfigurationSsl())
	certificates, _, err := ssl.CertificatesApi.GetCertificates(auth, nil)

	if err != nil {
		fmt.Printf("[ZonesApi] GetZones: %s", err)
	}

	for _, certificate := range certificates {
		fmt.Printf("%s: %s", certificate.CommonName, certificate.Status)
	}
}

func dnsExample(auth context.Context) {
	fmt.Println("========= DNS =========")
	dns := ionos.NewAPIClient(ionos.NewConfigurationDns())
	zones, _, err := dns.ZonesApi.GetZones(auth)
	if err != nil {
		fmt.Printf("[ZonesApi] GetZones: %s", err)
	}

	for _, zone := range zones {
		fmt.Printf("%s = %s\n", zone.Id, zone.Name)
	}
}
