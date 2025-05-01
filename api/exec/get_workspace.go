package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetWorkspaces() (*response.WorkspaceListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/workspace", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var workspaceListResponse response.WorkspaceListResponse
	err = json.Unmarshal(result, &workspaceListResponse)
	if err != nil {
		return nil, err
	}

	return &workspaceListResponse, nil
}

func GetWorkspace(id string) (*response.WorkspaceDetailResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/workspace/"+id, header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var workspaceResponse response.WorkspaceDetailResponse
	err = json.Unmarshal(result, &workspaceResponse)
	if err != nil {
		return nil, err
	}

	return &workspaceResponse, nil
}
