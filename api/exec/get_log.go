package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

type applicationLogs struct {
	Logs []string `json:"logs"`
}

func GetLog(workspaceId string, applicationId string) ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	response, err := api.SendGet("/"+workspaceId+"/application/"+applicationId+"/logs", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var applicationLogs applicationLogs
	err = json.Unmarshal(response, &applicationLogs)
	if err != nil {
		return nil, err
	}

	return applicationLogs.Logs, nil
}
