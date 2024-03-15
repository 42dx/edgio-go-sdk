package utils_test

import (
	"edgio/internal/utils"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockRoundTripper struct{}

func (m *MockRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, errors.New("mock error")
}

func TestGetHTTPJSONResultSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"key": "value"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

	httpClient := server.Client()
	request, _ := http.NewRequest(http.MethodGet, server.URL, nil)

	model, err := utils.GetHTTPJSONResult(httpClient, request, "token")
	require.NoError(t, err)

	assert.Equal(t, "value", model["key"])
}

func TestGetHTTPJSONResultNon200StatusCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))

	defer server.Close()

	httpClient := server.Client()
	request, _ := http.NewRequest(http.MethodGet, server.URL, nil)

	_, err := utils.GetHTTPJSONResult(httpClient, request, "token")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "[HTTP ERROR]: Status Code: 404")
}

func TestGetHTTPJSONResultDecodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`not a json`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

	httpClient := server.Client()
	request, _ := http.NewRequest(http.MethodGet, server.URL, nil)

	_, err := utils.GetHTTPJSONResult(httpClient, request, "token")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character")
}

func TestGetHTTPJSONResultHTTPClientError(t *testing.T) {
	mockTransport := &MockRoundTripper{}
	mockClient := &http.Client{Transport: mockTransport}
	request, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	token := "testToken"

	_, err := utils.GetHTTPJSONResult(mockClient, request, token)

	require.Error(t, err)
	require.Contains(t, err.Error(), "mock error")
}
