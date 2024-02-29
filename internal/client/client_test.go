package client

import (
	"edgio/common"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewMissingKey(t *testing.T) {
	creds := common.Creds{
		Key:     "",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: "http://example.com",
	}

	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := New(creds, config)

	assert.Equal(t, "edgio client key is missing", err.Error())
}

func TestNewMissingSecret(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "",
		Scopes:  "scopes",
		AuthUrl: "http://example.com",
	}

	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := New(creds, config)

	assert.Equal(t, "edgio client secret is missing", err.Error())
}

func TestNewUseDefaultScopes(t *testing.T) {
	defaultScopes := "app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, _ := io.ReadAll(req.Body)

		bodyStr := strings.ReplaceAll(string(body), "+", "%2B")
		params, _ := url.ParseQuery(bodyStr)
		assert.Equal(t, defaultScopes, params.Get("scope"))

		req.Body.Close()
	}))

	creds := common.Creds{
		Key:     "someKey",
		Secret:  "secret",
		AuthUrl: server.URL,
	}

	config := common.ClientConfig{
		Url:        server.URL,
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	New(creds, config)

	defer server.Close()
}

func TestNewMissingApiVersion(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: "http://example.com",
	}

	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := New(creds, config)

	assert.Equal(t, "edgio client config api version is missing", err.Error())
}

func TestNewMissingService(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: "http://example.com",
	}

	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "",
		Scope:      "scope",
	}

	_, err := New(creds, config)

	assert.Equal(t, "edgio client service is missing", err.Error())
}

func TestNewMissingScope(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: "http://example.com",
	}

	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "",
	}

	_, err := New(creds, config)

	assert.Equal(t, "edgio client scope is missing", err.Error())
}

func TestNewHappyPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))

	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: server.URL,
	}

	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	client, err := New(creds, config)

	assert.NoError(t, err)
	assert.Equal(t, "test_token", client.AccessToken)
	assert.Equal(t, config, client.Config)
	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestNewDefaultApiUrl(t *testing.T) {
	defaultEdgioApiUrl := "https://edgioapis.com"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))

	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: server.URL,
	}

	config := common.ClientConfig{
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	client, _ := New(creds, config)

	assert.Equal(t, defaultEdgioApiUrl, client.Config.Url)
	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestGetServiceUrlPathNotEmpty(t *testing.T) {
	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}
	client := Client{Config: config}
	params := common.UrlParams{Path: "path"}

	assert.Equal(t, "http://example.com/service/v1/scope/path", client.GetServiceUrl(params))
}

func TestGetServiceUrlPathEmpty(t *testing.T) {
	config := common.ClientConfig{
		Url:        "http://example.com",
		ApiVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}
	client := Client{Config: config}
	params := common.UrlParams{Path: ""}

	assert.Equal(t, "http://example.com/service/v1/scope", client.GetServiceUrl(params))
}
