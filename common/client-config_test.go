package common_test

// import (
// 	"edgio/common"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestMergeURL(t *testing.T) {
// 	c1 := common.ClientConfig{URL: "http://old-url.com"}
// 	c2 := common.ClientConfig{URL: "http://new-url.com"}

// 	result := c1.Merge(c2)

// 	assert.Equal(t, "http://new-url.com", result.URL)
// }

// func TestMergeAPIVersion(t *testing.T) {
// 	c1 := common.ClientConfig{APIVersion: "v1"}
// 	c2 := common.ClientConfig{APIVersion: "v2"}

// 	result := c1.Merge(c2)

// 	assert.Equal(t, "v2", result.APIVersion)
// }

// func TestMergeService(t *testing.T) {
// 	c1 := common.ClientConfig{Service: "OldService"}
// 	c2 := common.ClientConfig{Service: "NewService"}

// 	result := c1.Merge(c2)

// 	assert.Equal(t, "NewService", result.Service)
// }

// func TestMergeScope(t *testing.T) {
// 	c1 := common.ClientConfig{Scope: "OldScope"}
// 	c2 := common.ClientConfig{Scope: "NewScope"}

// 	result := c1.Merge(c2)

// 	assert.Equal(t, "NewScope", result.Scope)
// }

// func TestMergeOrgID(t *testing.T) {
// 	c1 := common.ClientConfig{OrgID: "OldOrgID"}
// 	c2 := common.ClientConfig{OrgID: "NewOrgID"}

// 	result := c1.Merge(c2)

// 	assert.Equal(t, "NewOrgID", result.OrgID)
// }
