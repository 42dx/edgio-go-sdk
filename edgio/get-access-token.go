package edgio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var result struct {
    AccessToken string `json:"access_token"`
}

var accessToken = ""

func GetAccessToken() {
    client     := &http.Client{}
    data       := url.Values{}

    if (credentials.Key == "" || credentials.Secret == "" || credentials.Scopes == "" || credentials.AuthUrl == "") {
        fmt.Println(errors.New("[AUTH ERROR]: Edgio client credentials are missing. Please set them using SetCreds() method").Error())
        os.Exit(1)
    }

    data.Set("client_id", credentials.Key)
    data.Set("client_secret", credentials.Secret)
    data.Set("grant_type", "client_credentials")
    data.Set("scope", credentials.Scopes)

    urlString, _  := url.QueryUnescape(data.Encode())
    request, _    := http.NewRequest(http.MethodPost, credentials.AuthUrl, strings.NewReader(urlString))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, err     := client.Do(request)
    if err != nil || resp.StatusCode != http.StatusOK {
        msg := []string{
            "[HTTP ERROR]: Status Code: ",
            strconv.Itoa(resp.StatusCode),
            " - ",
            http.StatusText(resp.StatusCode),
        }

        fmt.Println(errors.New(strings.Join(msg, "")).Error())
        os.Exit(1)
    }

    respData, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    err = json.Unmarshal([]byte(respData), &result)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    accessToken = result.AccessToken
}
