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

// evalCreds Validates and assigns default valued (if applicable) to the provided credentials.
// Mandatory params:
// common.Creds.Key
// common.Creds.Secret
// Optional params:
// common.Creds.Scopes
// common.Creds.AuthURL
// Returns a valid credentials and an error if any of the mandatory parameters are missing.
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

	if creds.AuthURL == "" {
		creds.AuthURL = "https://id.edgio.app/connect/token"
	}

	return creds, nil
}

// evalConfig Validates and assigns default valued (if applicable) to the provided configurations.
// Mandatory params:
// common.ClientConfig.APIVersion
// common.ClientConfig.Service
// common.ClientConfig.Scope
// Optional params:
// common.ClientConfig.URL
// Returns a valid client configuration and an error if any of the mandatory parameters are missing.
func evalConfig(config common.ClientConfig) (common.ClientConfig, error) {
	if config.URL == "" {
		config.URL = "https://edgioapis.com"
	}

	if config.APIVersion == "" {
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

// New Creates a base client configuration and connection to Edgio's REST API.
// The public packages under `edgio` namespace uses this client under the hood to perform their API calls.
// It validates and assing default valued (if applicable) to the provided credentials and configurations
// and returns a client instance with a valid access token.
// Mandatory params:
// common.Creds.Key
// common.Creds.Secret
// common.ClientConfig.APIVersion
// common.ClientConfig.Service
// common.ClientConfig.Scope
// Optional params:
// common.Creds.Scopes
// common.Creds.AuthURL
// common.ClientConfig.URL
// Returns a Client instance and an error if any of the mandatory parameters are missing.
func New(creds common.Creds, config common.ClientConfig) (Client, error) {
	accessToken := ""

	credentials, err := evalCreds(creds)
	if err != nil {
		return Client{}, err
	}

	configurations, err := evalConfig(config)
	if err != nil {
		return Client{}, err
	}

	if config.AccessToken != "" {
		accessToken = config.AccessToken
	} else {
		accessToken, err = token.GetAccessToken(credentials)
		if err != nil {
			return Client{}, err
		}
	}

	return Client{
		AccessToken: accessToken,
		Config:      configurations,
	}, nil
}

// GetServiceURL Returns the fully formatted Edgio REST API's url for the desired resource,
// identified by its `service`, `scope` and `apiVersion`.
// Mandatory params:
// Since this function inherits all its configuration from the created (by the `client.New` func) client,
// there are no mandatory parameters. All of them were already validated on the creation of the client.
// Optional params:
// common.URLParams.Path
// Returns the fully formatted Edgio REST API's url for the desired resource.
func (c Client) GetServiceURL(params common.URLParams) string {
	if params.Path != "" {
		params.Path = "/" + params.Path
	}

	return c.Config.URL + "/" + c.Config.Service + "/" + c.Config.APIVersion + "/" + c.Config.Scope + params.Path
}
