package main

import (
	"edgio/internal/client"
	"edgio/pkg/common"
	"edgio/pkg/org"
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
