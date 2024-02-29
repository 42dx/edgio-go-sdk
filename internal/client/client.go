package client

import (
	"edgio/common"
	"edgio/internal/token"
	"errors"
)

type Client struct {
	AccessToken string
	Config      common.ClientConfig
}

func evalCreds(creds common.Creds) (common.Creds, error) {
	if creds.Key == "" {
		return common.Creds{}, errors.New("edgio client key is missing")
	}

	if creds.Secret == "" {
		return common.Creds{}, errors.New("edgio client secret is missing")
	}

	if creds.Scopes == "" {
		creds.Scopes = "app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config"
	}

	if creds.AuthUrl == "" {
		creds.AuthUrl = "https://id.edgio.app/connect/token"
	}

	return creds, nil
}

func evalConfig(config common.ClientConfig) (common.ClientConfig, error) {
	if config.Url == "" {
		config.Url = "https://edgioapis.com"
	}

	if config.ApiVersion == "" {
		return common.ClientConfig{}, errors.New("edgio client config api version is missing")
	}

	if config.Service == "" {
		return common.ClientConfig{}, errors.New("edgio client service is missing")
	}

	if config.Scope == "" {
		return common.ClientConfig{}, errors.New("edgio client scope is missing")
	}

	return config, nil
}

func New(creds common.Creds, config common.ClientConfig) (Client, error) {
	credentials, err := evalCreds(creds)
	if err != nil {
		return Client{}, err
	}

	configurations, err := evalConfig(config)
	if err != nil {
		return Client{}, err
	}

	accessToken, err := token.GetAccessToken(credentials)
	if err != nil {
		return Client{}, err
	}

	return Client{
		AccessToken: accessToken,
		Config:      configurations,
	}, nil
}

func (c Client) GetServiceUrl(params common.UrlParams) string {
	if params.Path != "" {
		params.Path = "/" + params.Path
	}

	return c.Config.Url + "/" + c.Config.Service + "/" + c.Config.ApiVersion + "/" + c.Config.Scope + params.Path
}
