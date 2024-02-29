package org

import (
	"edgio/common"
	"edgio/internal/client"
	"edgio/internal/utils"
	"fmt"
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

func NewClient(params ClientParams) (OrgClientStruct, error) {
	fmt.Println(baseConfig)
	fmt.Println(params.Config)
	fmt.Println(baseConfig.Merge(params.Config))

	client, err := client.New(params.Credentials, baseConfig.Merge(params.Config))

	if err != nil {
		return OrgClientStruct{}, err
	}

	return OrgClientStruct{client}, nil
}

func (c OrgClientStruct) Get(params common.UrlParams) (getResultType, error) {
	httpClient := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, c.client.GetServiceUrl(common.UrlParams{Path: c.client.Config.OrgId}), nil)

	result, err := utils.GetHttpJsonResult(httpClient, request, c.client.AccessToken, &getResult)

	if err != nil {
		return getResultType{}, err
	}
	return *result.(*getResultType), nil
}
