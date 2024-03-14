package env_test

import (
	"edgio/common"
	"edgio/env"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc("/accounts/v0.1/environments", func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{
			"total_items": 2,
			"items": [
				{
					"id": "some-id",
					"name": "some-env-name",
					"legacy_account_number": "",
					"default_domain_name": "",
					"dns_domain_name": "",
					"can_members_deploy": true,
					"only_maintainers_can_deploy": true,
					"http_request_logging": true,
					"pci_compliance": true,
					"created_at": "2019-08-24T14:15:22Z",
					"updated_at": "2019-08-24T14:15:22Z"
				},
				{
					"id": "another-id",
					"name": "another-env-name",
					"legacy_account_number": "",
					"default_domain_name": "",
					"dns_domain_name": "",
					"can_members_deploy": false,
					"only_maintainers_can_deploy": false,
					"http_request_logging": false,
					"pci_compliance": false,
					"created_at": "2019-08-24T14:15:22Z",
					"updated_at": "2019-08-24T14:15:22Z"
				}
			]
		}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	result, _ := client.List("some-property-id")

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "some-env-name", result.Items[0].Name)
	assert.Equal(t, "another-env-name", result.Items[1].Name)
}

func TestListParseURLError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server.URL,
		},
		Config: common.ClientConfig{URL: ":"},
	}

	client, _ := env.NewClient(params)

	_, err := client.List("some-property-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse \":/accounts/v0.1/environments\": missing protocol scheme")
}

func TestListNewRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server.URL,
		},
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	_, err := client.List("\n")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid control character in URL")
}

func TestListGetHTTPJSONResultError(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc("/accounts/v0.1/environments", func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`error`))
		if err != nil {
			t.Fatal(err)
		}
	})

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := env.NewClient(params)
	_, err := client.List("some-property-id")

	require.Error(t, err)
	assert.Equal(t, "invalid character 'e' looking for beginning of value", err.Error())
}
