package common_test

import (
	"edgio/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableGetKey(t *testing.T) {
	v := common.Variable{Key: "testKey"}
	assert.Equal(t, "testKey", v.GetKey())
}

func TestVariableGetName(t *testing.T) {
	v := common.Variable{}
	assert.Equal(t, "", v.GetName())
}

func TestVariableGetSlug(t *testing.T) {
	v := common.Variable{}
	assert.Equal(t, "", v.GetSlug())
}

func TestPropertyGetKey(t *testing.T) {
	p := common.Property{}
	assert.Equal(t, "", p.GetKey())
}

func TestPropertyGetName(t *testing.T) {
	p := common.Property{}
	assert.Equal(t, "", p.GetName())
}

func TestPropertyGetSlug(t *testing.T) {
	p := common.Property{Slug: "testSlug"}
	assert.Equal(t, "testSlug", p.GetSlug())
}

func TestEnvGetKey(t *testing.T) {
	e := common.Env{}
	assert.Equal(t, "", e.GetKey())
}

func TestEnvGetName(t *testing.T) {
	e := common.Env{Name: "testName"}
	assert.Equal(t, "testName", e.GetName())
}

func TestEnvGetSlug(t *testing.T) {
	e := common.Env{}
	assert.Equal(t, "", e.GetSlug())
}
