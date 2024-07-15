package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	httpErr "github.com/dolong2/dcd-cli/api/err"
	"io"
	"net/http"
)

var baseUrl = "http://localhost:8081"

type apiErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendGet(targetUrl string, header map[string]string, param map[string]string) ([]byte, error) {
	result, err := sendHttpReq("GET", targetUrl, header, param, []byte(""))
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

func SendPost(targetUrl string, header map[string]string, body []byte) ([]byte, error) {
	result, err := sendHttpReq("POST", targetUrl, header, map[string]string{}, body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

func SendDelete(targetUrl string, header map[string]string, param map[string]string) ([]byte, error) {
	result, err := sendHttpReq("DELETE", targetUrl, header, param, []byte(""))
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

func SendPatch(targetUrl string, header map[string]string, body []byte) ([]byte, error) {
	result, err := sendHttpReq("PATCH", targetUrl, header, map[string]string{}, body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

func SendPut(targetUrl string, header map[string]string, body []byte) ([]byte, error) {
	result, err := sendHttpReq("PUT", targetUrl, header, map[string]string{}, body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

func sendHttpReq(method string, targetUrl string, header map[string]string, param map[string]string, body []byte) ([]byte, error) {
	if len(param) != 0 {
		targetUrl += "?"
	}
	for key, value := range param {
		targetUrl += fmt.Sprintf("%s=%s&", key, value)
	}

	httpClient := &http.Client{}
	request, err := http.NewRequest(method, baseUrl+targetUrl, bytes.NewBuffer(body))
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
		rawErrorResponse, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			return []byte(""), err
		}
		errorResponse := apiErrorResponse{}
		json.Unmarshal(rawErrorResponse, &errorResponse)
		return []byte(""), httpErr.NewHttpError(errorResponse.Status, errorResponse.Message)
	}

	result, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}
