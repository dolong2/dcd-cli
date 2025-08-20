package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetEnvList(workspaceId string) (*response.EnvListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/env", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var envListResponse response.EnvListResponse
	err = json.Unmarshal(result, &envListResponse)
	if err != nil {
		return nil, err
	}

	return &envListResponse, nil
}

func GetEnv(workspaceId string, envId string) (*response.EnvResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/env/"+envId, header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var envResponse response.EnvResponse
	err = json.Unmarshal(result, &envResponse)
	if err != nil {
		return nil, err
	}

	return &envResponse, nil
}
