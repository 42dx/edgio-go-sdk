package token_test

import (
	"edgio/common"
	"edgio/internal/token"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type brokenWriter struct{}

func (bw *brokenWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("broken writer")
}

type errorReader struct{}

func (e *errorReader) Read(_ []byte) (int, error) {
	return 0, errors.New("mocked read error")
}

func TestGetAccessTokenMissingKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		server.URL,
		httpmock.NewStringResponder(403, `{"access_token": "test_token"}`),
	)

	creds := common.Creds{
		Key:     "test_key",
		Scopes:  "test_scope",
		AuthURL: server.URL,
	}

	_, err := token.GetAccessToken(creds)
	require.Error(t, err)
	assert.Equal(t, "[AUTH ERROR]: Edgio client credentials are missing", err.Error())
}

func TestGetAccessTokenMissingSecret(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		server.URL,
		httpmock.NewStringResponder(403, `{"access_token": "test_token"}`),
	)

	creds := common.Creds{
		Secret:  "test_secret",
		Scopes:  "test_scope",
		AuthURL: server.URL,
	}

	_, err := token.GetAccessToken(creds)
	require.Error(t, err)
	assert.Equal(t, "[AUTH ERROR]: Edgio client credentials are missing", err.Error())
}

func TestGetAccessTokenInvalidAuthURL(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder(
		http.MethodPost,
		"https://id.edgio.app/invalid",
		httpmock.NewStringResponder(404, "Not Found"),
	)

	defer httpmock.DeactivateAndReset()

	creds := common.Creds{
		Key:     "test_key",
		Secret:  "test_secret",
		Scopes:  "test_scope",
		AuthURL: "https://id.edgio.app/invalid",
	}

	_, err := token.GetAccessToken(creds)
	require.Error(t, err)
	assert.Equal(t, "[HTTP ERROR]: Status Code: 404 - Not Found", err.Error())
}

func TestGetAccessTokenJsonUnmarshalError(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	server.Config.ErrorLog = log.New(&brokenWriter{}, "", 0)
	server.Start()
	defer server.Close()

	creds := common.Creds{
		Key:     "key",
		Secret:  "secret",
		Scopes:  "scopes",
		AuthURL: server.URL,
	}

	_, err := token.GetAccessToken(creds)
	require.Error(t, err)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestGetAccessToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		require.NoError(t, err)
	}))

	httpmock.Activate()

	httpmock.RegisterResponder(
		http.MethodPost,
		server.URL,
		httpmock.NewStringResponder(200, `{"access_token": "test_token"}`),
	)

	creds := common.Creds{
		Key:     "test_key",
		Secret:  "test_secret",
		Scopes:  "test_scope",
		AuthURL: server.URL,
	}

	token, _ := token.GetAccessToken(creds)

	assert.Equal(t, "test_token", token, "wrong test token")

	defer server.Close()
	defer httpmock.DeactivateAndReset()
}

func TestGetAccessTokenHttpError(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodPost,
		"https://id.edgio.app/connect/token",
		func(_ *http.Request) (*http.Response, error) {
			return nil, errors.New("mocked network error")
		},
	)

	creds := common.Creds{
		Key:     "testKey",
		Secret:  "testSecret",
		Scopes:  "testScopes",
		AuthURL: "https://id.edgio.app/connect/token",
	}

	_, err := token.GetAccessToken(creds)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mocked network error")
}

func TestGetAccessTokenReadBodyError(t *testing.T) {
	httpmock.Activate()

	defer httpmock.DeactivateAndReset()

	// Mocking an error while reading the response body
	httpmock.RegisterResponder(
		http.MethodPost,
		"https://id.edgio.app/connect/token",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "mocked response")
			resp.Body = io.NopCloser(&errorReader{})

			return resp, nil
		},
	)

	creds := common.Creds{
		Key:     "testKey",
		Secret:  "testSecret",
		Scopes:  "testScopes",
		AuthURL: "https://id.edgio.app/connect/token",
	}

	_, err := token.GetAccessToken(creds)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mocked read error")
}
