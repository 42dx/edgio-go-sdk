package property_test

import (
	"edgio/common"
	"edgio/property"
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
			OrgID: "some-org",
		},
	}

	client, err := property.NewClient(params)

	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestNewMissingOrdIDError(t *testing.T) {
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

	_, err := property.NewClient(params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OrgID is missing")
}

func TestNewClientError(t *testing.T) {
	params := common.ClientParams{
		Credentials: common.Creds{Key: "testKey"},
		Config:      common.ClientConfig{OrgID: "some-org"},
	}

	_, err := property.NewClient(params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "edgio client secret is missing")
}
