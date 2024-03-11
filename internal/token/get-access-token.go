package token

import (
	"edgio/common"
	"edgio/internal/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var Result struct {
	AccessToken string `json:"access_token"`
}

// GetAccessToken returns an access token for the provided credentials.
// Mandatory params:
// common.Creds.Key
// common.Creds.Secret
// Optional params:
// common.Creds.Scopes
// common.Creds.AuthURL
// Returns an access token and an error if any of the mandatory parameters are missing.
func GetAccessToken(credentials common.Creds) (string, error) {
	client := &http.Client{}

	if credentials.Key == "" || credentials.Secret == "" || credentials.Scopes == "" || credentials.AuthURL == "" {
		return "", errors.New("[AUTH ERROR]: Edgio client credentials are missing")
	}

	queryString := utils.ToQueryString(map[string]string{
		"client_id":     credentials.Key,
		"client_secret": credentials.Secret,
		"grant_type":    "client_credentials",
		"scope":         credentials.Scopes,
	})

	request, _ := http.NewRequest(http.MethodPost, credentials.AuthURL, strings.NewReader(queryString))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(request)
	if err != nil {
		return "", errors.New("[HTTP ERROR]: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := []string{
			"[HTTP ERROR]: Status Code: ",
			strconv.Itoa(resp.StatusCode),
			" - ",
			http.StatusText(resp.StatusCode),
		}

		return "", errors.New(strings.Join(msg, ""))
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = json.Unmarshal([]byte(respData), &Result)
	if err != nil {
		return "", err
	}

	return Result.AccessToken, nil
}
