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

type MockRoundTripper struct {
	RoundTripFunc func(_ *http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

	params := org.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server.URL,
		},
		Config: common.ClientConfig{
			URL:   "http://example.com",
			Scope: "scope",
		},
	}

	client, err := org.NewClient(params)

	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestGet(t *testing.T) {
	t.Parallel()

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

	params := org.ClientParams{
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

func TestNewClientError(t *testing.T) {
	t.Parallel()

	params := org.ClientParams{
		Credentials: common.Creds{Key: "testKey"},
		Config:      common.ClientConfig{},
	}

	_, err := org.NewClient(params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "edgio client secret is missing")
}

func TestGetError(t *testing.T) {
	t.Parallel()

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

	params := org.ClientParams{
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
