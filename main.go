package main

import (
	"edgio/common"
	"edgio/env"
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

	orgClient, _ := org.NewClient(common.ClientParams{Credentials: credentials})
	org, _ := orgClient.Get(common.URLParams{Path: os.Getenv("EDGIO_ORG_ID")})

	fmt.Println("Org ID: " + org.ID)

	propertyClient, err := property.NewClient(common.ClientParams{
		Credentials: credentials,
		Config:      common.ClientConfig{OrgID: org.ID, AccessToken: orgClient.Client.AccessToken},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	properties, _ := propertyClient.List()

	envClient, err := env.NewClient(common.ClientParams{
		Credentials: credentials,
		Config:      common.ClientConfig{AccessToken: orgClient.Client.AccessToken},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, property := range properties.Items {
		fmt.Println("Property: " + property.Slug)

		envs, err := envClient.List(property.ID)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, env := range envs.Items {
			fmt.Println("Env: " + env.Name)
		}
	}
	fmt.Println("main.go")
}
