package property

import (
	"edgio/common"
	"edgio/internal/client"
	"errors"
)

type ClientStruct struct {
	client client.Client
}

type Property struct {
	Id         string
	Slug       string
	Created_at string
	Updated_at string
}

type ClientParams struct {
	Credentials common.Creds
	Config      common.ClientConfig
}

var baseConfig = common.ClientConfig{
	APIVersion: "v0.1",
	Service:    "accounts",
	Scope:      "properties",
	OrgID:      "",
}

// NewClient returns a new client with the provided parameters.
// Mandatory params:
// common.Creds.Key
// common.Creds.Secret
// common.ClientConfig.OrgID
// Optional params:
// common.Creds.Scopes
// common.Creds.AuthURL
// common.ClientConfig.APIVersion
// common.ClientConfig.Service
// common.ClientConfig.Scope
// common.ClientConfig.URL
// Returns a new client and an error if any of the mandatory parameters are missing.
func NewClient(params ClientParams) (ClientStruct, error) {
	if params.Config.OrgID == "" {
		return ClientStruct{}, errors.New("OrgID is missing")
	}

	client, err := client.New(params.Credentials, baseConfig.Merge(params.Config))
	if err != nil {
		return ClientStruct{}, errors.New(err.Error())
	}

	return ClientStruct{client}, nil
}
