package client

import (
	"edgio/internal/token"
	"edgio/pkg/common"
	"errors"
	"fmt"
	"os"
)

type EdgioClientConfig struct {
    Url         string
    ApiVersion  string
    Service     string
    Scope       string
    OrgId       string
}

type EdgioClient struct {
    AccessToken string
    Config EdgioClientConfig
}

type UrlParams struct {
    Path    string
    Query   string
}

func (c EdgioClientConfig) Merge(other EdgioClientConfig) EdgioClientConfig {
    if other.ApiVersion != "" {
        c.ApiVersion = other.ApiVersion
    }
 
    if other.Service != "" {
        c.Service = other.Service
    }
 
    if other.Scope != "" {
        c.Scope = other.Scope
    }
 
    if other.OrgId != "" {
        c.OrgId = other.OrgId
    }
 
    return c
}

func evalCreds(creds common.Creds) common.Creds {
    if creds.Key == "" {
        fmt.Println(errors.New("edgio client key is missing"))
        os.Exit(1)
    }

    if creds.Secret == "" {
        fmt.Println(errors.New("edgio client secret is missing"))
        os.Exit(1)
    }

    if creds.Scopes == "" {
        creds.Scopes = "app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config"
    }

    if creds.AuthUrl == "" {
        creds.AuthUrl = "https://id.edgio.app/connect/token"
    }

    return creds
}

func evalConfig(config EdgioClientConfig) EdgioClientConfig {
    if config.Url == "" {
        config.Url = "https://edgioapis.com"
    }

    if config.ApiVersion == "" {
        fmt.Println(errors.New("edgio client config api version is missing"))
        os.Exit(1)
    }

    if config.Service == "" {
        fmt.Println(errors.New("edgio client service is missing"))
        os.Exit(1)
    }

    if config.Scope == "" {
        fmt.Println(errors.New("edgio client scope is missing"))
        os.Exit(1)
    }

    return config
}

func New(creds common.Creds, config EdgioClientConfig) EdgioClient {
    return EdgioClient{
        AccessToken: token.GetAccessToken(evalCreds(creds)),
        Config: evalConfig(config),
    }
}

func (c EdgioClient) GetServiceUrl(params UrlParams) string {
    if params.Path != "" {
        params.Path = "/" + params.Path
    }

    return c.Config.Url + "/" + c.Config.Service + "/" + c.Config.ApiVersion + "/" + c.Config.Scope + params.Path
}