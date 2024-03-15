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
	ID                       string    `mapstructure:"id"`
	Name                     string    `mapstructure:"name"`
	LegacyAccNumber          string    `mapstructure:"legacy_account_number"`
	DefaultDomainName        string    `mapstructure:"default_domain_name"`
	DNSDomainName            string    `mapstructure:"dns_domain_name"`
	CanMembersDeploy         bool      `mapstructure:"can_members_deploy"`
	OnlyMaintainersCanDeploy bool      `mapstructure:"only_maintainers_can_deploy"`
	HTTPRequestLogging       bool      `mapstructure:"http_request_logging"`
	PciCompliance            bool      `mapstructure:"pci_compliance"`
	CreatedAt                time.Time `mapstructure:"create_at"`
	UpdatedAt                time.Time `mapstructure:"updated_at"`
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
