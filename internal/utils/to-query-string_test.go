package utils_test

import (
	"edgio/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToQueryString(t *testing.T) {
	params := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	queryString := utils.ToQueryString(params)

	assert.Equal(t, "key1=value1&key2=value2", queryString)
}
