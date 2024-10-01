package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResult struct {
	Output []string `json:"result"`
}

func ExecCommand(workspaceId string, applicationId string, command string) ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	commandRequest := CommandRequest{Command: command}
	requestBody, err := json.Marshal(commandRequest)
	if err != nil {
		return nil, err
	}

	response, err := api.SendPost("/"+workspaceId+"/application/"+applicationId+"/exec", header, map[string]string{}, requestBody)
	if err != nil {
		return nil, err
	}

	var commandResult CommandResult
	err = json.Unmarshal(response, &commandResult)
	if err != nil {
		return nil, err
	}

	return commandResult.Output, nil
}
