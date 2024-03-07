package utils

import (
	"errors"
	"net/url"
)

func ToQueryString(params map[string]string) (string, error) {
	queryParams := url.Values{}

	for key, value := range params {
		queryParams.Set(key, value)
	}

	urlString, err := url.QueryUnescape(queryParams.Encode())
	if err != nil {
		return "", errors.New(err.Error())
	}

	return urlString, nil
}
