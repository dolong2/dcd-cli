package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

type WorkspaceListResponse struct {
	List []WorkspaceResponse `json:"list"`
}

type WorkspaceResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetWorkspace() (*WorkspaceListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	response, err := api.SendGet("/workspace", header)
	if err != nil {
		return nil, err
	}

	var workspaceListResponse WorkspaceListResponse
	err = json.Unmarshal(response, &workspaceListResponse)
	if err != nil {
		return nil, err
	}

	return &workspaceListResponse, nil
}
