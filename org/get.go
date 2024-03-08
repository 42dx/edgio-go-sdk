package org

import (
	"edgio/common"
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

// Get returns the organization details.
// Mandatory params:
// common.URLParams.OrgID
// Returns the organization details and an error if any of the mandatory parameters are missing.
func (c ClientStruct) Get(params common.URLParams) (GetResultType, error) {
	httpClient := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, c.Client.GetServiceURL(params), nil)

	err := utils.GetHTTPJSONResult(httpClient, request, c.Client.AccessToken, &getResult)
	if err != nil {
		return GetResultType{}, errors.New(err.Error())
	}

	return getResult, nil
}
