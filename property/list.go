package property

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
	Items []Property `json:"items"`
}

var listResult = listResultType{}

// List lists the properties from a given Organization.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all properties
// until actual pagination is implemented.
// Returns a list of properties for a given Organization, or an error if anything goes wrong.
func (c ClientStruct) List() (listResultType, error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	rawQueryString := utils.ToQueryString(
		map[string]string{"organization_id": c.Config.OrgID, "page_size": "100"},
	)

	parsedURL.RawQuery = rawQueryString

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	properties, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return listResultType{}, errors.New(err.Error())
	}

	mapstructure.Decode(properties, &listResult)

	return listResult, nil
}
