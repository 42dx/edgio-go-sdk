package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeUrl(t *testing.T) {
	c1 := ClientConfig{Url: "http://old-url.com"}
	c2 := ClientConfig{Url: "http://new-url.com"}

	result := c1.Merge(c2)

	assert.Equal(t, "http://new-url.com", result.Url)
}

func TestMergeApiVersion(t *testing.T) {
	c1 := ClientConfig{ApiVersion: "v1"}
	c2 := ClientConfig{ApiVersion: "v2"}

	result := c1.Merge(c2)

	assert.Equal(t, "v2", result.ApiVersion)
}

func TestMergeService(t *testing.T) {
	c1 := ClientConfig{Service: "OldService"}
	c2 := ClientConfig{Service: "NewService"}

	result := c1.Merge(c2)

	assert.Equal(t, "NewService", result.Service)
}

func TestMergeScope(t *testing.T) {
	c1 := ClientConfig{Scope: "OldScope"}
	c2 := ClientConfig{Scope: "NewScope"}

	result := c1.Merge(c2)

	assert.Equal(t, "NewScope", result.Scope)
}

func TestMergeOrgId(t *testing.T) {
	c1 := ClientConfig{OrgId: "OldOrgId"}
	c2 := ClientConfig{OrgId: "NewOrgId"}

	result := c1.Merge(c2)

	assert.Equal(t, "NewOrgId", result.OrgId)
}
