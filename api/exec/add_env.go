package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"strings"
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

	_, err = api.SendPost("/"+workspaceId+"/application/"+applicationId+"/env", header, map[string]string{}, requestJson)
	if err != nil {
		return err
	}

	return nil
}

func AddEnvWithLabels(workspaceId string, labels []string, key string, value string) error {
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

	params := make(map[string]string)
	params["labels"] = strings.Join(labels, ",")

	_, err = api.SendPost("/"+workspaceId+"/application/env", header, params, requestJson)
	if err != nil {
		return err
	}

	return nil
}
