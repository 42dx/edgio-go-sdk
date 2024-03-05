package org

import (
	"edgio/common"
	"edgio/internal/client"
	"edgio/internal/utils"
	"errors"
	"net/http"
)

type GetResultType struct {
	ID   string
	Name string
}

var getResult = GetResultType{
	ID:   `json:"id"`,
	Name: `json:"name"`,
}

type ClientStruct struct {
	client client.Client
}

type ClientParams struct {
	Credentials common.Creds
	Config      common.ClientConfig
}

var baseConfig = common.ClientConfig{
	APIVersion: "v0.1",
	Service:    "accounts",
	Scope:      "organizations",
	OrgID:      "",
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
func NewClient(params ClientParams) (ClientStruct, error) {
	client, err := client.New(params.Credentials, baseConfig.Merge(params.Config))
	if err != nil {
		return ClientStruct{}, errors.New(err.Error())
	}

	return ClientStruct{client}, nil
}

// Get returns the organization details.
// Mandatory params:
// common.URLParams.OrgID
// Returns the organization details and an error if any of the mandatory parameters are missing.
func (c ClientStruct) Get(params common.URLParams) (GetResultType, error) {
	httpClient := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, c.client.GetServiceURL(params), nil)

	result, err := utils.GetHTTPJSONResult(httpClient, request, c.client.AccessToken, &getResult)
	if err != nil {
		return GetResultType{}, errors.New(err.Error())
	}

	return *result.(*GetResultType), nil
}
