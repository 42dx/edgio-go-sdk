package org_test

import (
	"edgio/common"
	"edgio/org"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
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

	mux.HandleFunc("/accounts/v0.1/organizations/some-id", func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"id": "some-id", "name": "some-name"}`))
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

	client, _ := org.NewClient(params)
	result, _ := client.Get(common.URLParams{Path: "some-id"})

	assert.Equal(t, "some-id", result.ID)
	assert.Equal(t, "some-name", result.Name)
}

func TestGetError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/accounts/v0.1/organizations/some-id", func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`not a json`))
		if err != nil {
			t.Fatal(err)
		}
	})

	server := httptest.NewServer(mux)

	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server2.Close()

	params := common.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL},
	}

	orgClient, _ := org.NewClient(params)

	_, err := orgClient.Get(common.URLParams{Path: "some-id"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character")
}
