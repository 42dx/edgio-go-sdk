package token

import (
	"edgio/common"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
	if err == nil || !strings.Contains(err.Error(), expected) {
		t.Fatalf("Expected error to contain '%s' but got '%s'", expected, err.Error())
	}
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
