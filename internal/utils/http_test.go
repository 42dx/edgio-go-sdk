package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHttpJsonResultSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(`{"key": "value"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

	httpClient := server.Client()
	request, _ := http.NewRequest("GET", server.URL, nil)
	model := make(map[string]string)

	result, err := GetHttpJsonResult(httpClient, request, "token", &model)
	assert.NoError(t, err)

	resultModel := result.(*map[string]string)
	assert.Equal(t, "value", (*resultModel)["key"])
}

func TestGetHttpJsonResultNon200StatusCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))

	defer server.Close()

	httpClient := server.Client()
	request, _ := http.NewRequest("GET", server.URL, nil)
	model := make(map[string]string)

	_, err := GetHttpJsonResult(httpClient, request, "token", &model)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[HTTP ERROR]: Status Code: 404")
}

func TestGetHttpJsonResultDecodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(`not a json`))
		if err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

	httpClient := server.Client()
	request, _ := http.NewRequest("GET", server.URL, nil)
	model := make(map[string]string)

	_, err := GetHttpJsonResult(httpClient, request, "token", &model)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character")
}
