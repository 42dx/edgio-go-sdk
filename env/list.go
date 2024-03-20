package env

import (
	"edgio/common"
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

type ListResult common.BaseListResultType[common.Env]
type FilterParams struct {
	PropertyID string
	Name       string
}

var envListResult ListResult

// List Lists the environments for a given Property.
// Edgio's list page size was defaulted to 100 for now,
// which is the highest value. The idea is to return all environments
// until actual pagination is implemented.
// Returns a list of environments for a given Property, or an error if anything goes wrong.
func (c ClientStruct) List(propertyID string) (ListResult, error) {
	httpClient := &http.Client{}
	serviceURL := c.GetServiceURL(common.URLParams{})

	parsedURL, err := url.Parse(serviceURL)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	parsedURL.RawQuery = utils.ToQueryString(
		map[string]string{"page_size": "100", "property_id": propertyID},
	)

	request, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	envsJSONmap, err := utils.GetHTTPJSONResult(httpClient, request, c.AccessToken)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	err = mapstructure.Decode(envsJSONmap, &envListResult)
	if err != nil {
		return ListResult{}, errors.New(err.Error())
	}

	return envListResult, nil
}

func (c ClientStruct) FilterList(params FilterParams) (common.FilteredListResultType[common.Env], error) {
	if params.PropertyID == "" {
		return common.FilteredListResultType[common.Env]{}, errors.New("PropertyID is required")
	}

	fullEnvList, err := c.List(params.PropertyID)

	if params.Name == "" {
		return common.FilteredListResultType[common.Env]{
			BaseListResultType: common.BaseListResultType[common.Env]{
				Total: fullEnvList.Total,
				Items: fullEnvList.Items,
			},
			FilteredTotal: fullEnvList.Total,
		}, nil
	}

	if err != nil {
		return common.FilteredListResultType[common.Env]{}, errors.New(err.Error())
	}

	filteredProperties := utils.FilterList[common.Env](
		utils.FilterListParams[common.Env]{Needle: params.Name, Haystack: fullEnvList.Items},
	)

	return common.FilteredListResultType[common.Env]{
		BaseListResultType: common.BaseListResultType[common.Env]{
			Total: fullEnvList.Total,
			Items: filteredProperties,
		},
		FilteredTotal: len(filteredProperties),
	}, nil
}
