package main

import (
	"context"

	_ "github.com/stevenklar/ionos-hosting-api/pkg/client"
)

func main() {
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key:    "APIKEY",
		Prefix: "Bearer",
	})
	r, err := client.Service.Operation(auth, args)
}
