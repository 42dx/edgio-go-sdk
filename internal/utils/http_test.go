package utils_test

// import (
// 	"edgio/internal/utils"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestGetHTTPJSONResultSuccess(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
// 		_, err := rw.Write([]byte(`{"key": "value"}`))
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	}))

// 	defer server.Close()

// 	httpClient := server.Client()
// 	request, _ := http.NewRequest(http.MethodGet, server.URL, nil)
// 	model := make(map[string]string)

// 	result, err := utils.GetHTTPJSONResult(httpClient, request, "token", &model)
// 	require.NoError(t, err)

// 	resultModel := result.(*map[string]string)
// 	assert.Equal(t, "value", (*resultModel)["key"])
// }

// func TestGetHTTPJSONResultNon200StatusCode(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
// 		rw.WriteHeader(http.StatusNotFound)
// 	}))

// 	defer server.Close()

// 	httpClient := server.Client()
// 	request, _ := http.NewRequest(http.MethodGet, server.URL, nil)
// 	model := make(map[string]string)

// 	_, err := utils.GetHTTPJSONResult(httpClient, request, "token", &model)
// 	require.Error(t, err)
// 	assert.Contains(t, err.Error(), "[HTTP ERROR]: Status Code: 404")
// }

// func TestGetHTTPJSONResultDecodeError(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
// 		_, err := rw.Write([]byte(`not a json`))
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	}))

// 	defer server.Close()

// 	httpClient := server.Client()
// 	request, _ := http.NewRequest(http.MethodGet, server.URL, nil)
// 	model := make(map[string]string)

// 	_, err := utils.GetHTTPJSONResult(httpClient, request, "token", &model)
// 	require.Error(t, err)
// 	assert.Contains(t, err.Error(), "invalid character")
// }
