package variable_test

import (
	"edgio/common"
	"edgio/variable"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const envVarURL = "/config/v0.1/environment-variables"
const tokenReturn = `{"access_token": "test_token"}`
const variablesResponse = `{
    "total_items": 2,
    "items": [
        {
            "id": "some-id",
            "key": "some-env-var-key",
            "value": "some-value",
            "secret": true,
            "create_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        },
        {
            "id": "another-id",
            "key": "another-env-var-key",
            "value": "another-value",
            "secret": false,
            "create_at": "2019-08-24T14:15:22Z",
            "updated_at": "2019-08-24T14:15:22Z"
        }
    ]
}`

func TestList(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(tokenReturn))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc(envVarURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(variablesResponse))
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

	client, _ := variable.NewClient(params)
	result, _ := client.List("some-environment-id")

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "some-env-var-key", result.Items[0].Key)
	assert.Equal(t, "another-env-var-key", result.Items[1].Key)
}

func TestListParseURLError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(tokenReturn))
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
		Config: common.ClientConfig{URL: ":"},
	}

	client, _ := variable.NewClient(params)

	_, err := client.List("some-environment-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse \":/config/v0.1/environment-variables\": missing protocol scheme")
}

func TestListNewRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(tokenReturn))
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := variable.NewClient(params)
	_, err := client.List("\n")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid control character in URL")
}

func TestListGetHTTPJSONResultError(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(tokenReturn))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc(envVarURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`error`))
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

	client, _ := variable.NewClient(params)
	_, err := client.List("some-environment-id")

	require.Error(t, err)
	assert.Equal(t, "invalid character 'e' looking for beginning of value", err.Error())
}

func TestListMapstructureDecodeError(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(tokenReturn))
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

	mux.HandleFunc(envVarURL, func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"items": "invalid"}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	client, _ := variable.NewClient(params)
	_, err := client.List("some-env-id")

	require.Error(t, err)
	assert.Equal(t, "1 error(s) decoding:\n\n* 'items': source data must be an array or slice, got string", err.Error())
}

func TestFilterList(t *testing.T) {
	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(tokenReturn))
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
	}

	t.Run("returns error when List fails", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		params.Config.URL = server.URL

		defer server.Close()

		mux.HandleFunc(envVarURL, func(rw http.ResponseWriter, _ *http.Request) {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		})

		client, _ := variable.NewClient(params)
		_, err := client.FilterList(variable.FilterParams{EnvID: "some-environment-id"})

		assert.Equal(t, "[HTTP ERROR]: Status Code: 500 - Internal Server Error", err.Error())
		require.Error(t, err)
	})

	t.Run("returns error when EnvID is not provided", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		params.Config.URL = server.URL

		client, _ := variable.NewClient(params)
		_, err := client.FilterList(variable.FilterParams{})

		assert.Error(t, err)
		assert.Equal(t, "EnvID is required", err.Error())
	})

	t.Run("returns full list when no key is provided", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		params.Config.URL = server.URL

		mux.HandleFunc(envVarURL, func(rw http.ResponseWriter, _ *http.Request) {
			rw.Write([]byte(variablesResponse))
		})

		client, _ := variable.NewClient(params)
		result, err := client.FilterList(variable.FilterParams{EnvID: "some-environment-id"})

		assert.NoError(t, err)
		assert.Equal(t, 2, result.Total)
		assert.Equal(t, 2, result.FilteredTotal)
	})

	t.Run("returns filtered list when key is provided", func(t *testing.T) {
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		params.Config.URL = server.URL

		mux.HandleFunc(envVarURL, func(rw http.ResponseWriter, _ *http.Request) {
			rw.Write([]byte(variablesResponse))
		})

		client, _ := variable.NewClient(params)
		result, err := client.FilterList(variable.FilterParams{EnvID: "some-environment-id", Key: "another-env-var-key"})

		assert.NoError(t, err)
		assert.Equal(t, 2, result.Total)
		assert.Equal(t, 1, result.FilteredTotal)
		assert.Equal(t, "another-env-var-key", result.Items[0].Key)
	})
}
