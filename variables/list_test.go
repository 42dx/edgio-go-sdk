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

var envVarsResponse = `{
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
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc("/config/v0.1/environment-variables", func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(envVarsResponse))
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

	client, _ := envVar.NewClient(params)
	result, _ := client.List("some-environment-id")

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "some-env-var-key", result.Items[0].Key)
	assert.Equal(t, "another-env-var-key", result.Items[1].Key)
}

func TestListParseURLError(t *testing.T) {
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
		Config: common.ClientConfig{URL: ":"},
	}

	client, _ := envVar.NewClient(params)

	_, err := client.List("some-environment-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse \":/config/v0.1/environment-variables\": missing protocol scheme")
}

func TestListNewRequestError(t *testing.T) {
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
		Config: common.ClientConfig{URL: server.URL},
	}

	client, _ := envVar.NewClient(params)
	_, err := client.List("\n")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid control character in URL")
}

func TestListGetHTTPJSONResultError(t *testing.T) {
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

	mux.HandleFunc("/config/v0.1/environment-variables", func(rw http.ResponseWriter, _ *http.Request) {
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

	client, _ := envVar.NewClient(params)
	_, err := client.List("some-environment-id")

	require.Error(t, err)
	assert.Equal(t, "invalid character 'e' looking for beginning of value", err.Error())
}
