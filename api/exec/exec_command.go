package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func ExecCommand(workspaceId string, applicationId string, command string) ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	commandRequest := request.CommandRequest{Command: command}
	requestBody, err := json.Marshal(commandRequest)
	if err != nil {
		return nil, err
	}

	result, err := api.SendPost("/"+workspaceId+"/application/"+applicationId+"/exec", header, map[string]string{}, requestBody)
	if err != nil {
		return nil, err
	}

	var commandResult response.CommandResult
	err = json.Unmarshal(result, &commandResult)
	if err != nil {
		return nil, err
	}

	return commandResult.Output, nil
}
