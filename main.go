package main

import (
	"edgio/common"
	"edgio/org"
	"edgio/property"
	"fmt"
	"os"
)

func main() {
	fmt.Println("main.go")
	credentials := common.Creds{
		Key:    os.Getenv("EDGIO_API_CLIENT_ID"),
		Secret: os.Getenv("EDGIO_API_CLIENT_SECRET"),
	}

	orgClient, _ := org.NewClient(org.ClientParams{Credentials: credentials})
	org, _ := orgClient.Get(common.URLParams{Path: os.Getenv("EDGIO_ORG_ID")})

	fmt.Println("Org ID: " + org.ID)

	propertyClient, err := property.NewClient(property.ClientParams{
		Credentials: credentials,
		Config:      common.ClientConfig{OrgID: org.ID},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	properties, _ := propertyClient.List()

	fmt.Println(properties.Total)
	for _, item := range properties.Items {
		fmt.Printf("Property slug: " + item.Slug + "\n")
	}
	fmt.Println("main.go")
}
