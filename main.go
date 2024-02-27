package main

import (
	"42dx/pkg/client"
	"42dx/pkg/common"
	"42dx/pkg/org"
	"fmt"
	"os"
)

func main() {
    credentials := common.Creds{
        Key: os.Getenv("EDGIO_CLIENT_KEY"),
        Secret: os.Getenv("EDGIO_CLIENT_SECRET"),
    }

    orgClient := org.NewClient(org.ClientParams{
        Credentials: credentials,
        Config: client.EdgioClientConfig{ OrgId: os.Getenv("ORG_ID") },
    })

    fmt.Println(orgClient.Get(client.UrlParams{}))
}
