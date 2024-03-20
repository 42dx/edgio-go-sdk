package utils

import (
	"edgio/common"
	"strings"
)

type Filterable interface {
	common.Searchable
	common.Property | common.Env | common.Variable
}

type FilterListParams[T Filterable] struct {
	Needle   string
	Haystack []T
}

func FilterList[T Filterable](params FilterListParams[T]) []T {
	result := []T{}

	for _, item := range params.Haystack {
		if strings.Contains(item.GetKey(), params.Needle) ||
			strings.Contains(item.GetName(), params.Needle) ||
			strings.Contains(item.GetSlug(), params.Needle) {
			result = append(result, item)
		}
	}

	return result
}
