package main

import (
	"edgio/common"
	"edgio/org"
	"fmt"
	"os"
)

func main() {
	credentials := common.Creds{
		Key:    os.Getenv("EDGIO_CLIENT_KEY"),
		Secret: os.Getenv("EDGIO_CLIENT_SECRET"),
	}

	orgClient, _ := org.NewClient(org.ClientParams{
		Credentials: credentials,
		Config:      common.EdgioClientConfig{OrgId: os.Getenv("ORG_ID")},
	})

	fmt.Println(orgClient.Get(common.UrlParams{}))
}
