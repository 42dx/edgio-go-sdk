package variable

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
	Items []Variable `json:"items"`
}

var listResult = listResultType{}

// List Lists the environment variables for a given Environment.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all environment variables
// until actual pagination is implemented.
// Returns a list of environment variables for a given Environment, or an error if anything goes wrong.
func (c ClientStruct) List(environmentID string) (listResultType, error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	rawQueryString := utils.ToQueryString(
		map[string]string{"page_size": "100", "environment_id": environmentID},
	)

	parsedURL.RawQuery = rawQueryString

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	variables, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	mapstructure.Decode(variables, &listResult)

	return listResult, nil
}
