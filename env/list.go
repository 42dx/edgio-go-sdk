package env

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

type listResultType struct {
	common.BaseListResultType
	Items []Env `json:"items"`
}

var listResult = listResultType{}

// List Lists the environments for a given Property.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all environments
// until actual pagination is implemented.
// Returns a list of environments for a given Property, or an error if anything goes wrong.
func (c ClientStruct) List(propertyID string) (listResultType, error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	rawQueryString := utils.ToQueryString(
		map[string]string{"page_size": "100", "property_id": propertyID},
	)

	parsedURL.RawQuery = rawQueryString

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	envs, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	mapstructure.Decode(envs, &listResult)

	return listResult, nil
}
