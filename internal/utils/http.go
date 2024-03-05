package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// GetHttpJsonResult returns the result of an HTTP request in JSON format.
// Mandatory params:
// httpClient *http.Client
// request *http.Request
// token string
// model interface{}
// Returns the result of an HTTP request in JSON format and an error if the request fails.
func GetHttpJsonResult(httpClient *http.Client, request *http.Request, token string, model interface{}) (interface{}, error) {
	request.Header.Add("Authorization", token)

	resp, err := httpClient.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		msg := []string{
			"[HTTP ERROR]: Status Code: ",
			strconv.Itoa(resp.StatusCode),
			" - ",
			http.StatusText(resp.StatusCode),
		}

		return nil, errors.New(strings.Join(msg, ""))
	}

	err = json.NewDecoder(resp.Body).Decode(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
