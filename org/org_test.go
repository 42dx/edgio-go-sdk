package org

import (
	"edgio/common"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))

	params := ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthUrl: server.URL,
		},
		Config: common.ClientConfig{
			Url:   "http://example.com",
			Scope: "scope",
		},
	}

	client, err := NewClient(params)

	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/accounts/v0.1/organizations/some-id", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"id": "some-id", "name": "some-name"}`))
	})
	server := httptest.NewServer(mux)
	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))

	defer server.Close()

	params := ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthUrl: server2.URL,
		},
		Config: common.ClientConfig{
			Url:   server.URL,
			OrgId: "some-id",
		},
	}

	client, _ := NewClient(params)
	result, err := client.Get(common.UrlParams{})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "some-id", result.Id)
	assert.Equal(t, "some-name", result.Name)
	defer server.Close()
}
