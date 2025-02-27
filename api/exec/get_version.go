package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

type resourceVersion struct {
	VersionList []string `json:"version"`
}

func GetVersion(resourceType string) ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	response, err := api.SendGet("/application/"+resourceType+"/version", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var resourceVersion resourceVersion
	err = json.Unmarshal(response, &resourceVersion)
	if err != nil {
		return nil, err
	}

	return resourceVersion.VersionList, nil
}
