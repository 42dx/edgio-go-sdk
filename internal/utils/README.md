# internal/utils

This package package holds some utility functions that are used to outsource some common logic from other packages to avoid repetition/ease testing.

## `utils.GetHttpJsonResult(httpClient *http.Client, request *http.Request, token string, model interface{}) (interface{}, error)`

This function has mainly three related goals: 1- Process http requests; 2- treat HTTP errors in a standardized way, and; 3- Process and decode returned json data from the endpoints.

```go
type resultType struct { Id string }
var result = resultType{ Id: `json:"id"` }

httpClient := &http.Client{}
request, _ := http.NewRequest(http.MethodGet, client.GetServiceUrl(common.UrlParams{ Path: "some-org-id" }), nil)

result, err := utils.GetHttpJsonResult(httpClient, request, "some-access-token", &result)
```

### `utils.GetHttpJsonResult` Mandatory Params

- `httpClient *http.Client`: Http client instance from the `net/http` package.
- `request *http.Request`: The request configs (result of `net/http.NewRequest` function).
- `token string`: The Edgio API access token.
- `model interface{}`: Struct of the expected data response.

### `utils.GetHttpJsonResult` Optional Params & Default Values

There is no optional parameters for that function

<p align="right"><em><a href="../../#utils">back to the main README</a></em></p>
