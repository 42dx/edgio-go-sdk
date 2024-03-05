package client_test

import (
	"edgio/common"
	"edgio/internal/client"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMissingKey(t *testing.T) {
	creds := common.Creds{
		Key:     "",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: "http://example.com",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := client.New(creds, config)
	require.Error(t, err)
	assert.Equal(t, "edgio client key is missing", err.Error(), "Error message is not as expected")
}

func TestNewMissingSecret(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "",
		Scopes:  "scopes",
		AuthURL: "http://example.com",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := client.New(creds, config)

	require.Error(t, err)
	assert.Equal(t, "edgio client secret is missing", err.Error(), "Error message is not as expected")
}

func TestNewUseDefaultScopes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, req *http.Request) {
		defaultScopes := "app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config"
		body, _ := io.ReadAll(req.Body)
		bodyStr := strings.ReplaceAll(string(body), "+", "%2B")
		params, _ := url.ParseQuery(bodyStr)
		assert.Equal(t, defaultScopes, params.Get("scope"))

		req.Body.Close()
	}))
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	creds := common.Creds{
		Key:     "someKey",
		Secret:  "secret",
		AuthURL: server2.URL,
	}

	config := common.ClientConfig{
		URL:        server.URL,
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := client.New(creds, config)
	require.NoError(t, err)
}

func TestNewMissingAPIVersion(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: "http://example.com",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "",
		Service:    "service",
		Scope:      "scope",
	}

	_, err := client.New(creds, config)

	assert.Equal(t, "edgio client config api version is missing", err.Error())
}

func TestNewMissingService(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: "http://example.com",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "",
		Scope:      "scope",
	}

	_, err := client.New(creds, config)
	require.Error(t, err)
	assert.Equal(t, "edgio client service is missing", err.Error())
}

func TestGetTokenErr(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodPost, "https://id.edgio.app/connect/token", httpmock.NewStringResponder(403, "forbidden"))

	defer httpmock.DeactivateAndReset()

	creds := common.Creds{
		Key:    "testKey",
		Secret: "testSecret",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "org",
	}

	_, err := client.New(creds, config)

	require.Error(t, err)
	assert.Equal(t, "[HTTP ERROR]: Status Code: 403 - Forbidden", err.Error())
}

func TestDefaultAuthURL(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodPost, "https://id.edgio.app/connect/token", httpmock.NewStringResponder(200, `{"access_token": "test_token"}`))

	defer httpmock.DeactivateAndReset()

	creds := common.Creds{
		Key:    "testKey",
		Secret: "testSecret",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "org",
	}

	_, err := client.New(creds, config)

	require.NoError(t, err)
}

func TestNewMissingScope(t *testing.T) {
	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: "http://example.com",
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "",
	}

	_, err := client.New(creds, config)
	require.Error(t, err)
	assert.Equal(t, "edgio client scope is missing", err.Error())
}

func TestNewHappyPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()
	defer httpmock.DeactivateAndReset()

	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: server.URL,
	}

	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}

	client, err := client.New(creds, config)

	require.NoError(t, err)
	assert.Equal(t, "test_token", client.AccessToken)
	assert.Equal(t, config, client.Config)
}

func TestNewDefaultApiURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()
	defer httpmock.DeactivateAndReset()

	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: server.URL,
	}
	config := common.ClientConfig{
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}
	client, _ := client.New(creds, config)
	defaultEdgioAPI := "https://edgioapis.com"

	assert.Equal(t, defaultEdgioAPI, client.Config.URL)
}

func TestGetServiceURLPathNotEmpty(t *testing.T) {
	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}
	client := client.Client{Config: config}
	params := common.URLParams{Path: "path"}

	assert.Equal(t, "http://example.com/service/v1/scope/path", client.GetServiceURL(params))
}

func TestGetServiceURLPathEmpty(t *testing.T) {
	config := common.ClientConfig{
		URL:        "http://example.com",
		APIVersion: "v1",
		Service:    "service",
		Scope:      "scope",
	}
	client := client.Client{Config: config}
	params := common.URLParams{Path: ""}

	assert.Equal(t, "http://example.com/service/v1/scope", client.GetServiceURL(params))
}
