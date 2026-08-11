package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	httpErr "github.com/dolong2/dcd-cli/api/err"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

var baseUrl string

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

func SendPost(targetUrl string, header map[string]string, param map[string]string, body []byte) ([]byte, error) {
	result, err := sendHttpReq("POST", targetUrl, header, param, body)
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

func SendPatch(targetUrl string, header map[string]string, param map[string]string, body []byte) ([]byte, error) {
	result, err := sendHttpReq("PATCH", targetUrl, header, param, body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

func SendPut(targetUrl string, header map[string]string, param map[string]string, body []byte) ([]byte, error) {
	result, err := sendHttpReq("PUT", targetUrl, header, param, body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}

// SendPostMultipart 파일 업로드와 같이 multipart/form-data로 파일을 전송해야 하는 요청에 사용
func SendPostMultipart(targetUrl string, header map[string]string, param map[string]string, fileFieldName string, fileName string, fileContent io.Reader) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fileFieldName, fileName)
	if err != nil {
		return []byte(""), err
	}
	if _, err := io.Copy(part, fileContent); err != nil {
		return []byte(""), err
	}
	if err := writer.Close(); err != nil {
		return []byte(""), err
	}

	request, err := http.NewRequest("POST", baseUrl+withQuery(targetUrl, param), body)
	if err != nil {
		return []byte(""), err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	for key, value := range header {
		request.Header.Set(key, value)
	}

	return doRequest(request)
}

func sendHttpReq(method string, targetUrl string, header map[string]string, param map[string]string, body []byte) ([]byte, error) {
	request, err := http.NewRequest(method, baseUrl+withQuery(targetUrl, param), bytes.NewBuffer(body))
	if err != nil {
		return []byte(""), err
	}

	request.Header.Set("Content-Type", "application/json")
	for key, value := range header {
		request.Header.Set(key, value)
	}

	return doRequest(request)
}

func withQuery(targetUrl string, param map[string]string) string {
	if len(param) == 0 {
		return targetUrl
	}

	query := url.Values{}
	for key, value := range param {
		query.Add(key, value)
	}

	return fmt.Sprintf("%s?%s", targetUrl, query.Encode())
}

func doRequest(request *http.Request) ([]byte, error) {
	httpClient := &http.Client{}
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
		if len(rawErrorResponse) == 0 {
			return []byte(""), httpErr.NewHttpError(httpResponse.StatusCode, "알 수 없는 에러")
		}
		errorResponse := apiErrorResponse{}
		err = json.Unmarshal(rawErrorResponse, &errorResponse)
		if err != nil {
			return []byte(""), err
		}
		return []byte(""), httpErr.NewHttpError(errorResponse.Status, errorResponse.Message)
	}

	result, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return []byte(""), err
	}

	return result, nil
}
