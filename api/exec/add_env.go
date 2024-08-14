package exec

import (
	"encoding/json"
	"fmt"
	"github.com/dolong2/dcd-cli/api"
)

type envRequest struct {
	EnvList map[string]string `json:"envList"`
}

func AddEnv(workspaceId string, applicationId string, key string, value string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	body := map[string]string{}
	body[key] = value
	envReq := envRequest{EnvList: body}
	requestJson, err := json.Marshal(envReq)

	fmt.Println(envReq)
	fmt.Println(string(requestJson))

	_, err = api.SendPost("/"+workspaceId+"/application/"+applicationId+"/env", header, map[string]string{}, requestJson)
	if err != nil {
		return err
	}

	return nil
}
