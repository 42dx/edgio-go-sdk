package utils_test

import (
	"edgio/common"
	"edgio/internal/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterList(t *testing.T) {
	propertyHaystack := []common.Property{
		{Slug: "slug1"},
		{Slug: "slug2"},
		{Slug: "slug3"},
	}

	envHaystack := []common.Env{
		{Name: "name1"},
		{Name: "name2"},
		{Name: "name3"},
	}

	variableHaystack := []common.Variable{
		{Key: "key1"},
		{Key: "key2"},
		{Key: "key3"},
	}

	t.Run("returns items that match the needle (with 'Slug')", func(t *testing.T) {
		params := utils.FilterListParams[common.Property]{Needle: "1", Haystack: propertyHaystack}
		result := utils.FilterList[common.Property](params)

		assert.Len(t, result, 1)
		assert.Equal(t, "slug1", result[0].GetSlug())
	})

	t.Run("returns items that match the needle (with 'Name')", func(t *testing.T) {
		params := utils.FilterListParams[common.Env]{Needle: "1", Haystack: envHaystack}
		result := utils.FilterList[common.Env](params)

		assert.Len(t, result, 1)
		assert.Equal(t, "name1", result[0].GetName())
	})

	t.Run("returns items that match the needle (with 'Key')", func(t *testing.T) {
		params := utils.FilterListParams[common.Variable]{Needle: "1", Haystack: variableHaystack}
		result := utils.FilterList[common.Variable](params)

		assert.Len(t, result, 1)
		assert.Equal(t, "key1", result[0].GetKey())
	})

	t.Run("returns no items when none match the needle", func(t *testing.T) {
		params := utils.FilterListParams[common.Property]{Needle: "4", Haystack: propertyHaystack}
		result := utils.FilterList[common.Property](params)

		assert.Len(t, result, 0)
	})

	t.Run("filter is case sensitive", func(t *testing.T) {
		params := utils.FilterListParams[common.Property]{Needle: strings.ToUpper("slug1"), Haystack: propertyHaystack}
		result := utils.FilterList[common.Property](params)

		assert.Len(t, result, 0)
	})
}
