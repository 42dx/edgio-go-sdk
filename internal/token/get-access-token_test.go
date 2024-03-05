package token

import (
	"edgio/common"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

type brokenWriter struct{}

func (bw *brokenWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("broken writer")
}

func TestGetAccessTokenMissingKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))
	httpmock.Activate()

	httpmock.RegisterResponder("POST", server.URL,
		httpmock.NewStringResponder(403, `{"access_token": "test_token"}`))

	creds := common.Creds{
		Key:     "test_key",
		Scopes:  "test_scope",
		AuthUrl: server.URL,
	}

	_, err := GetAccessToken(creds)

	assert.Equal(t, "[AUTH ERROR]: Edgio client credentials are missing", err.Error())

	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestGetAccessTokenMissingSecret(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))
	httpmock.Activate()

	httpmock.RegisterResponder("POST", server.URL,
		httpmock.NewStringResponder(403, `{"access_token": "test_token"}`))

	creds := common.Creds{
		Secret:  "test_secret",
		Scopes:  "test_scope",
		AuthUrl: server.URL,
	}

	_, err := GetAccessToken(creds)

	assert.Equal(t, "[AUTH ERROR]: Edgio client credentials are missing", err.Error())

	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestGetAccessTokenInvalidAuthUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))
	httpmock.Activate()

	httpmock.RegisterResponder("POST", server.URL,
		httpmock.NewStringResponder(404, "Not Found"))

	creds := common.Creds{
		Key:     "test_key",
		Secret:  "test_secret",
		Scopes:  "test_scope",
		AuthUrl: server.URL,
	}

	_, err := GetAccessToken(creds)

	assert.Equal(t, "[HTTP ERROR]: Status Code: 404 - Not Found", err.Error())

	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestGetAccessTokenJsonUnmarshalError(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	server.Config.ErrorLog = log.New(&brokenWriter{}, "", 0)
	server.Start()
	defer server.Close()

	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthUrl: server.URL,
	}

	_, err := GetAccessToken(creds)

	expected := "unexpected end of JSON input"
	assert.Equal(t, expected, err.Error())
}

func TestGetAccessToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"access_token": "test_token"}`))
	}))

	httpmock.Activate()
	httpmock.RegisterResponder("POST", server.URL,
		httpmock.NewStringResponder(200, `{"access_token": "test_token"}`))

	creds := common.Creds{
		Key:     "test_key",
		Secret:  "test_secret",
		Scopes:  "test_scope",
		AuthUrl: server.URL,
	}

	token, _ := GetAccessToken(creds)

	assert.Equal(t, token, "test_token", "wrong test token")

	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestGetAccessToken_HttpError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mocking a network error
	httpmock.RegisterResponder("POST", "https://id.edgio.app/connect/token",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("mocked network error")
		},
	)

	creds := common.Creds{
		Key:     "testKey",
		Secret:  "testSecret",
		Scopes:  "testScopes",
		AuthUrl: "https://id.edgio.app/connect/token",
	}

	_, err := GetAccessToken(creds)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mocked network error")
}

func TestGetAccessToken_ReadBodyError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mocking an error while reading the response body
	httpmock.RegisterResponder("POST", "https://id.edgio.app/connect/token",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "mocked response")
			resp.Body = io.NopCloser(&errorReader{})
			return resp, nil
		},
	)

	creds := common.Creds{
		Key:     "testKey",
		Secret:  "testSecret",
		Scopes:  "testScopes",
		AuthUrl: "https://id.edgio.app/connect/token",
	}

	_, err := GetAccessToken(creds)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mocked read error")
}

type errorReader struct{}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mocked read error")
}
