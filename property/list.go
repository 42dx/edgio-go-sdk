package property

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

type ListResult common.BaseListResultType[common.Property]
type FilterParams struct {
	Slug string
}

var propertyListResult ListResult

// List lists the properties from a given Organization.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all properties
// until actual pagination is implemented.
// Returns a list of properties for a given Organization, or an error if anything goes wrong.
func (c ClientStruct) List() (ListResult, error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	parsedURL.RawQuery = utils.ToQueryString(
		map[string]string{"organization_id": c.Config.OrgID, "page_size": "100"},
	)

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	propertiesJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(propertiesJSONmap, &propertyListResult)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	return propertyListResult, nil
}

func (c ClientStruct) FilterList(params FilterParams) (common.FilteredListResultType[common.Property], error) {
	fullPropertyList, err := c.List()
	if err != nil {
		return common.FilteredListResultType[common.Property]{}, errors.New(err.Error())
	}

	if params.Slug == "" {
		return common.FilteredListResultType[common.Property]{
			BaseListResultType: common.BaseListResultType[common.Property]{
				Total: fullPropertyList.Total,
				Items: fullPropertyList.Items,
			},
			FilteredTotal: fullPropertyList.Total,
		}, nil
	}

	filteredProperties := utils.FilterList[common.Property](utils.FilterListParams[common.Property]{Needle: params.Slug, Haystack: fullPropertyList.Items})

	return common.FilteredListResultType[common.Property]{
		BaseListResultType: common.BaseListResultType[common.Property]{
			Total: fullPropertyList.Total,
			Items: filteredProperties,
		},
		FilteredTotal: len(filteredProperties),
	}, nil
}
