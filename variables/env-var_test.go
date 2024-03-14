package variables_test

import (
	common "edgio/common"
	envVar "edgio/variables"
	http "net/http"
	httptest "net/http/httptest"
	testing "testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

type MockRoundTripper struct {
	RoundTripFunc func(_ *http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestNewClient(t *testing.T) {
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
		Config: common.ClientConfig{
			URL:   "http://example.com",
			Scope: "scope",
		},
	}

	client, err := envVar.NewClient(params)

	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestNewClientError(t *testing.T) {
	params := common.ClientParams{
		Credentials: common.Creds{Key: "testKey"},
		Config:      common.ClientConfig{},
	}

	_, err := envVar.NewClient(params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "edgio client secret is missing")
}
