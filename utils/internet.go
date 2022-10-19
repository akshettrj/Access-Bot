package utils

import (
	"net/http"
	"net/http/cookiejar"
)

func HTTPClientWithCookieJar() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, nil
	}
	return &http.Client{Jar: jar}, nil
}

func HTTPGetRequest(request *http.Request, client *http.Client) (int, *http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	response, err := client.Do(request)
	if err != nil {
		return 0, response, err
	}
	return response.StatusCode, response, err
}
