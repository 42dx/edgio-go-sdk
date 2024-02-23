package main

import (
	"42dx/edgio"
	"os"
)

func main() {
    creds := edgio.EdgioCreds{
        Key: os.Getenv("EDGIO_CLIENT_KEY"),
        Secret: os.Getenv("EDGIO_CLIENT_SECRET"),
    }

    edgio.SetCreds(creds)
    edgio.GetAccessToken()
}
