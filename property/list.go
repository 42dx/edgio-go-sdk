package property

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"
)

type ListResultType struct {
	common.BaseListResultType
	Items []Property `json:"items"`
}

var PropertyListResult = ListResultType{}

// List list of properties for a given Organization.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all properties
// until actual pagination is implemented.
// Returns a list of properties for a given Organization or an error if anything goes wrong.
func (c ClientStruct) List() (ListResultType, error) {
	httpClient := &http.Client{}
	queryStringMap := map[string]string{"organization_id": c.client.Config.OrgID}
	serviceURL := c.client.GetServiceURL(common.URLParams{Query: map[string]string{"page_size": "100"}})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return ListResultType{}, errors.New(err.Error())
	}

	rawQueryString := utils.ToQueryString(queryStringMap)

	parsedURL.RawQuery = rawQueryString

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return ListResultType{}, errors.New(err.Error())
	}

	err = utils.GetHTTPJSONResult(httpClient, request, c.client.AccessToken, &PropertyListResult)
	if err != nil {
		return ListResultType{}, errors.New(err.Error())
	}

	return PropertyListResult, nil
}
