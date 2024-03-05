package main

import (
	"edgio/common"
	"edgio/org"
	"fmt"
	"os"
)

func main() {
	credentials := common.Creds{
		Key:    os.Getenv("EDGIO_API_CLIENT_ID"),
		Secret: os.Getenv("EDGIO_API_CLIENT_SECRET"),
	}

	orgClient, _ := org.NewClient(org.ClientParams{Credentials: credentials})

	fmt.Println(orgClient.Get(common.URLParams{Path: os.Getenv("EDGIO_ORG_ID")}))
}
