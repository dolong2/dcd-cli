package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetVolumeList(workspaceId string) (*response.VolumeListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/volume", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var volumeListResponse response.VolumeListResponse
	err = json.Unmarshal(result, &volumeListResponse)
	if err != nil {
		return nil, err
	}

	return &volumeListResponse, nil
}
