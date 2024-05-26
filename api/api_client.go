package api

import (
	"bytes"
	"io"
	"net/http"
)

func SendPost(targetUrl string, header map[string]string, body []byte) ([]byte, error) {
	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer(body))
	if err != nil {
		return []byte(""), err
	}

	request.Header.Set("Content-Type", "application/json")
	for key, value := range header {
		request.Header.Set(key, value)
	}

	httpResponse, err := httpClient.Do(request)
	if err != nil {
		return []byte(""), err
	}
	defer httpResponse.Body.Close()

	result, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}
