package token

// import (
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/jarcoal/httpmock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAccessTokenMissingKey(t *testing.T) {
//     server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//         rw.Write([]byte(`{"access_token": "test_token"}`))
//     }))
//     httpmock.Activate()

//     httpmock.RegisterResponder("POST", server.URL,
//         httpmock.NewStringResponder(200, `{"access_token": "test_token"}`))

//     creds := EdgioCreds{
//         Key: "test_key",
//         Scopes: "test_scope",
//         AuthUrl: server.URL,
//     }

//     _, err := GetAccessToken(creds)

//     assert.Equal(t, errors.New("Edgio client secret is missing"), err)

//     defer server.Close()
//     defer httpmock.DeactivateAndReset()
// }

// func TestGetAccessTokenMissingSecret(t *testing.T) {
//     server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//         rw.Write([]byte(`{"access_token": "test_token"}`))
//     }))
//     httpmock.Activate()

//     httpmock.RegisterResponder("POST", server.URL,
//         httpmock.NewStringResponder(200, `{"access_token": "test_token"}`))

//     creds := EdgioCreds{
//         Secret: "test_secret",
//         Scopes: "test_scope",
//         AuthUrl: server.URL,
//     }

//     _, err := GetAccessToken(creds)

//     assert.Equal(t, errors.New("Edgio client key is missing"), err)

//     defer server.Close()
//     defer httpmock.DeactivateAndReset()
// }

// func TestGetAccessTokenInvalidAuthUrl(t *testing.T) {
//     server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//         rw.Write([]byte(`{"access_token": "test_token"}`))
//     }))
//     httpmock.Activate()

//     httpmock.RegisterResponder("POST", server.URL,
//         httpmock.NewStringResponder(404, "Not Found"))

//     creds := EdgioCreds{
//         Key:    "test_key",
//         Secret: "test_secret",
//         Scopes: "test_scope",
//         AuthUrl: server.URL,
//     }

//     _, err := GetAccessToken(creds)

//     assert.Equal(t, errors.New("[HTTP ERROR]: Status Code: 404 - Not Found"), err)

//     defer server.Close()
//     defer httpmock.DeactivateAndReset()
// }

// func TestGetAccessToken(t *testing.T) {
//     server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//         rw.Write([]byte(`{"access_token": "test_token"}`))
//     }))

//     httpmock.Activate()
//     httpmock.RegisterResponder("POST", server.URL,
//     httpmock.NewStringResponder(200, `{"access_token": "test_token"}`))

//     creds := EdgioCreds{
//         Key:    "test_key",
//         Secret: "test_secret",
//         Scopes: "test_scope",
//         AuthUrl: server.URL,
//     }

//     token, _ := GetAccessToken(creds)

//     assert.Equal(t, token, "test_token", "wrong test token")

//     defer server.Close()
//     defer httpmock.DeactivateAndReset()
// }
