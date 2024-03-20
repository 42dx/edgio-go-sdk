package property

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

var propertyListResult common.BaseListResultType[common.Property]

// List lists the properties from a given Organization.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all properties
// until actual pagination is implemented.
// Returns a list of properties for a given Organization, or an error if anything goes wrong.
func (c ClientStruct) List() (common.BaseListResultType[common.Property], error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return common.BaseListResultType[common.Property]{}, errors.New(err.Error())
	}

	rawQueryString := utils.ToQueryString(
		map[string]string{"organization_id": c.Config.OrgID, "page_size": "100"},
	)

	parsedURL.RawQuery = rawQueryString

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return common.BaseListResultType[common.Property]{}, errors.New(err.Error())
	}

	propertiesJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return common.BaseListResultType[common.Property]{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(propertiesJSONmap, &propertyListResult)
	if err != nil {
		return common.BaseListResultType[common.Property]{}, errors.New(err.Error())
	}

	return propertyListResult, nil
}

func (c ClientStruct) FilterList(slug string) (common.FilteredListResultType[common.Property], error) {
	fullPropertyList, err := c.List()
	if err != nil {
		return common.FilteredListResultType[common.Property]{}, errors.New(err.Error())
	}

	filteredProperties := utils.FilterList[common.Property](utils.FilterListParams[common.Property]{Needle: slug, Haystack: fullPropertyList.Items})

	return common.FilteredListResultType[common.Property]{
		BaseListResultType: common.BaseListResultType[common.Property]{
			Total: fullPropertyList.Total,
			Items: filteredProperties,
		},
		FilteredTotal: len(filteredProperties),
	}, nil
}
