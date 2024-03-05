package org

import (
	"edgio/common"
	"edgio/internal/client"
	"edgio/internal/utils"
	"net/http"
)

type getResultType struct {
	Id   string
	Name string
}

var getResult = getResultType{
	Id:   `json:"id"`,
	Name: `json:"name"`,
}

type OrgClientStruct struct {
	client client.Client
}

type ClientParams struct {
	Credentials common.Creds
	Config      common.ClientConfig
}

var baseConfig = common.ClientConfig{
	ApiVersion: "v0.1",
	Service:    "accounts",
	Scope:      "organizations",
	OrgId:      "",
}

// NewClient returns a new client with the provided parameters.
// Mandatory params:
// common.Creds.Key
// common.Creds.Secret
// Optional params:
// common.Creds.Scopes
// common.Creds.AuthUrl
// common.ClientConfig.ApiVersion
// common.ClientConfig.Service
// common.ClientConfig.Scope
// common.ClientConfig.Url
// Returns a new client and an error if any of the mandatory parameters are missing.
func NewClient(params ClientParams) (OrgClientStruct, error) {
	client, err := client.New(params.Credentials, baseConfig.Merge(params.Config))
	if err != nil {
		return OrgClientStruct{}, err
	}

	return OrgClientStruct{client}, nil
}

// Get returns the organization details.
// Mandatory params:
// common.UrlParams.OrgId
// Returns the organization details and an error if any of the mandatory parameters are missing.
func (c OrgClientStruct) Get(params common.UrlParams) (getResultType, error) {
	httpClient := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, c.client.GetServiceUrl(params), nil)

	result, err := utils.GetHttpJsonResult(httpClient, request, c.client.AccessToken, &getResult)
	if err != nil {
		return getResultType{}, err
	}

	return *result.(*getResultType), nil
}
