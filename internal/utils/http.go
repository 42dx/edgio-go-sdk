package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// GetHTTPJSONResult returns the result of an HTTP request in JSON format.
// Mandatory params:
// httpClient *http.Client
// request *http.Request
// token string
// Returns the result of an HTTP request in JSON format and an error if the request fails.
func GetHTTPJSONResult(httpClient *http.Client, request *http.Request, token string) (map[string]any, error) {
	var mappedJSONResult map[string]any

	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("content-type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := []string{
			"[HTTP ERROR]: Status Code: ",
			strconv.Itoa(resp.StatusCode),
			" - ",
			http.StatusText(resp.StatusCode),
		}

		return nil, errors.New(strings.Join(msg, ""))
	}

	err = json.NewDecoder(resp.Body).Decode(&mappedJSONResult)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return mappedJSONResult, nil
}
