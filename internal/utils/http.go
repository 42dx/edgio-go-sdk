package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetHttpJsonResult(httpClient *http.Client, request *http.Request, token string, model interface{}) interface{} {
	request.Header.Add("Authorization", token)

	resp, err	:= httpClient.Do(request)
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

	err = json.NewDecoder(resp.Body).Decode(&model)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	return model
}