package env

import (
	"edgio/common"
	"edgio/internal/client"
	"errors"
	"time"
)

type ClientStruct struct {
	*client.Client
}

type Env struct {
	ID                       string
	Name                     string
	LegacyAccNumber          string
	DefaultDomainName        string
	DNSDomainName            string
	CanMembersDeploy         bool
	OnlyMaintainersCanDeploy bool
	HTTPRequestLogging       bool
	PciCompliance            bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

var baseConfig = common.ClientConfig{
	APIVersion: "v0.1",
	Service:    "accounts",
	Scope:      "environments",
}

// NewClient returns a new client with the provided parameters.
// Mandatory params:
// common.Creds.Key
// common.Creds.Secret
// Optional params:
// common.Creds.Scopes
// common.Creds.AuthURL
// common.ClientConfig.APIVersion
// common.ClientConfig.Service
// common.ClientConfig.Scope
// common.ClientConfig.URL
// Returns a new client and an error if any of the mandatory parameters are missing.
func NewClient(params common.ClientParams) (ClientStruct, error) {
	client, err := client.New(params.Credentials, baseConfig.Merge(params.Config))
	if err != nil {
		return ClientStruct{}, errors.New(err.Error())
	}

	return ClientStruct{&client}, nil
}
