package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
	"strings"
)

func GetApplications(workspaceId string) (*response.ApplicationListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/application", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var applicationListResponse response.ApplicationListResponse
	err = json.Unmarshal(result, &applicationListResponse)
	if err != nil {
		return nil, err
	}

	return &applicationListResponse, nil
}

func GetApplicationsByLabels(workspaceId string, labels []string) (*response.ApplicationListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["labels"] = strings.Join(labels, ",")

	result, err := api.SendGet("/"+workspaceId+"/application", header, param)
	if err != nil {
		return nil, err
	}

	var applicationListResponse response.ApplicationListResponse
	err = json.Unmarshal(result, &applicationListResponse)
	if err != nil {
		return nil, err
	}

	return &applicationListResponse, nil
}

func GetApplication(workspaceId string, applicationId string) (*response.ApplicationResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/application/"+applicationId, header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var applicationResponse response.ApplicationResponse
	err = json.Unmarshal(result, &applicationResponse)
	if err != nil {
		return nil, err
	}

	return &applicationResponse, nil
}
