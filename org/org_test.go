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
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

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
		_, err := rw.Write([]byte(`{"id": "some-id", "name": "some-name"}`))
		if err != nil {
			t.Fatal(err)
		}
	})
	server := httptest.NewServer(mux)
	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()
	defer server2.Close()

	params := ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthUrl: server2.URL,
		},
		Config: common.ClientConfig{Url: server.URL},
	}

	client, _ := NewClient(params)
	result, _ := client.Get(common.UrlParams{Path: "some-id"})

	assert.Equal(t, "some-id", result.Id)
	assert.Equal(t, "some-name", result.Name)
}

type MockRoundTripper struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestNewClientError(t *testing.T) {
	params := ClientParams{
		Credentials: common.Creds{Key: "testKey"},
		Config:      common.ClientConfig{},
	}

	_, err := NewClient(params)
	if err != nil {
		t.Fatal(err)
	}

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "edgio client secret is missing")
}

func TestGetError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/accounts/v0.1/organizations/some-id", func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(`not a json`))
		if err != nil {
			t.Fatal(err)
		}
	})
	server := httptest.NewServer(mux)
	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()
	defer server2.Close()

	params := ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthUrl: server2.URL,
		},
		Config: common.ClientConfig{Url: server.URL},
	}

	orgClient, _ := NewClient(params)

	_, err := orgClient.Get(common.UrlParams{Path: "some-id"})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character")
}
