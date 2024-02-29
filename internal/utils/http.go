package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetHttpJsonResult(httpClient *http.Client, request *http.Request, token string, model interface{}) (interface{}, error) {
	request.Header.Add("Authorization", token)

	fmt.Println(request)
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
