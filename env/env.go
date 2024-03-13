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
	ID                       string    `json:"id"`
	Name                     string    `json:"name"`
	LegacyAccNumber          string    `json:"legacy_account_number"`
	DefaultDomainName        string    `json:"default_domain_name"`
	DnsDomainName            string    `json:"dns_domain_name"`
	CanMembersDeploy         bool      `json:"can_members_deploy"`
	OnlyMaintainersCanDeploy bool      `json:"only_maintainers_can_deploy"`
	HttpRequestLogging       bool      `json:"http_request_logging"`
	PciCompliance            bool      `json:"pci_compliance"`
	CreatedAt                time.Time `json:"create_at"`
	UpdatedAt                time.Time `json:"updated_at"`
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
