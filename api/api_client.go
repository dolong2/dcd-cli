package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

var baseUrl = "http://localhost:8081"

func SendPost(targetUrl string, header map[string]string, body []byte) ([]byte, error) {
	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", baseUrl+targetUrl, bytes.NewBuffer(body))
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

	// 200번대 응답코드가 아닐때 에러
	if httpResponse.StatusCode/100 != 2 {
		return []byte(""), errors.New("response status code is not 2xx")
	}

	result, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}
