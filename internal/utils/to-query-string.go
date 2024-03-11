package utils

import "net/url"

func ToQueryString(params map[string]string) string {
	queryParams := url.Values{}

	for key, value := range params {
		queryParams.Set(key, value)
	}

	urlString, _ := url.QueryUnescape(queryParams.Encode())

	return urlString
}
