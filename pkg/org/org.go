package org

import (
	"42dx/internal/utils"
	"42dx/pkg/client"
	"42dx/pkg/common"
	"net/http"
)

type getResultType struct {
	Id			string 
	Name		string 
}

var getResult = getResultType{
	Id:		`json:"id"`,
	Name:	`json:"name"`,
}

type OrgClientStruct struct {
	client client.EdgioClient
}

type ClientParams struct {
	Credentials common.Creds
	Config client.EdgioClientConfig
}

var baseConfig = client.EdgioClientConfig{
	ApiVersion: "v0.1",
	Service:	"accounts",
	Scope:		"organizations",
	OrgId:		"",
}

func NewClient(params ClientParams) OrgClientStruct {
	return OrgClientStruct{
		client: client.New(params.Credentials, baseConfig.Merge(params.Config)),
	}
}

func (c OrgClientStruct) Get(params client.UrlParams) getResultType {
	httpClient	:= &http.Client{}
	request, _	:= http.NewRequest(http.MethodGet, c.client.GetServiceUrl(client.UrlParams{ Path: c.client.Config.OrgId }), nil)

	utils.GetHttpJsonResult(httpClient, request, c.client.AccessToken, &getResult)

	return getResult
}
