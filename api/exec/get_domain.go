package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetDomains(workspaceId string) (*response.DomainListResponse, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/domain", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var domainListResponse response.DomainListResponse
	err = json.Unmarshal(result, &domainListResponse)
	if err != nil {
		return nil, err
	}

	return &domainListResponse, nil
}
