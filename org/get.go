package org

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"

	"github.com/mitchellh/mapstructure"
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
// Returns the organization details, or an error if any of the mandatory parameters are missing
// or any underlying process goes wrong.
func (c ClientStruct) Get(params common.URLParams) (GetResultType, error) {
	httpClient := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, c.Client.GetServiceURL(params), nil)

	JSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.Client.AccessToken)

	if err != nil {
		return GetResultType{}, errors.New(err.Error())
	}

	mapstructure.Decode(JSONmap, &getResult)

	return getResult, nil
}
