package edgio

import (
	"errors"
	"fmt"
	"os"
)

type EdgioCreds struct {
    Key     string
    Secret  string
    Scopes  string
    AuthUrl string
}

var credentials EdgioCreds

func evalParams(creds EdgioCreds) (EdgioCreds, error) {
    if creds.Key == "" {
        return EdgioCreds{Key: "", Secret: "", Scopes: "", AuthUrl: ""}, errors.New("edgio client key is missing")
    }

    if creds.Secret == "" {
        return EdgioCreds{Key: "", Secret: "", Scopes: "", AuthUrl: ""}, errors.New("edgio client secret is missing")
    }

    if creds.Scopes == "" {
        creds.Scopes = "app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config"
    }

    if creds.AuthUrl == "" {
        creds.AuthUrl = "https://id.edgio.app/connect/token"
    }

    return creds, nil
}

func SetCreds(creds EdgioCreds) {
    creds, err := evalParams(creds)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    credentials = creds
}
