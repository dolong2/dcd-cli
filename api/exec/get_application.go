package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

type ApplicationResponse struct {
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	ApplicationType string            `json:"applicationType"`
	GithubUrl       string            `json:"githubUrl"`
	Env             map[string]string `json:"env"`
	Port            int               `json:"port"`
	ExternalPort    int               `json:"externalPort"`
	Version         string            `json:"version"`
	Status          string            `json:"status"`
}

type ApplicationListResponse struct {
	Applications []ApplicationResponse `json:"list"`
}

func GetApplications(workspaceId string) (*ApplicationListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	params := map[string]string{"workspaceId": workspaceId}
	response, err := api.SendGet("/application", header, params)
	if err != nil {
		return nil, err
	}

	var applicationListResponse ApplicationListResponse
	err = json.Unmarshal(response, &applicationListResponse)
	if err != nil {
		return nil, err
	}

	return &applicationListResponse, nil
}
