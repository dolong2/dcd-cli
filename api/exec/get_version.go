package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetVersion(resourceType string) ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/application/"+resourceType+"/version", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var resourceVersion response.ResourceVersion
	err = json.Unmarshal(result, &resourceVersion)
	if err != nil {
		return nil, err
	}

	return resourceVersion.VersionList, nil
}
