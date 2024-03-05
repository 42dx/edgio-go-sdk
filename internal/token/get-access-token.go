package token

import (
	"edgio/common"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
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
	data := url.Values{}

	if credentials.Key == "" || credentials.Secret == "" || credentials.Scopes == "" || credentials.AuthURL == "" {
		return "", errors.New("[AUTH ERROR]: Edgio client credentials are missing")
	}

	data.Set("client_id", credentials.Key)
	data.Set("client_secret", credentials.Secret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", credentials.Scopes)

	urlString, _ := url.QueryUnescape(data.Encode())
	request, _ := http.NewRequest(http.MethodPost, credentials.AuthURL, strings.NewReader(urlString))
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
