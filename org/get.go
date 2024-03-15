package org

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

var getResult Org

// Get returns the organization details.
// Mandatory params:
// common.URLParams.OrgID
// Returns the organization details, or an error if any of the mandatory parameters are missing
// or any underlying process goes wrong.
func (c ClientStruct) Get(params common.URLParams) (Org, error) {
	httpClient := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, c.Client.GetServiceURL(params), nil)

	orgJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.Client.AccessToken)
	if err != nil {
		return Org{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(orgJSONmap, &getResult)
	if err != nil {
		return Org{}, errors.New(err.Error())
	}

	return getResult, nil
}
