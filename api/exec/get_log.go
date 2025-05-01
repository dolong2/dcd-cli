package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetLog(workspaceId string, applicationId string) ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/"+workspaceId+"/application/"+applicationId+"/logs", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var applicationLogs response.ApplicationLogs
	err = json.Unmarshal(result, &applicationLogs)
	if err != nil {
		return nil, err
	}

	return applicationLogs.Logs, nil
}
