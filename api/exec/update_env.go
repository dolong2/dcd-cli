package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"strings"
)

type updateEnvRequest struct {
	NewValue string `json:"newValue"`
}

func UpdateEnv(workspaceId string, applicationId string, key string, value string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	updateEnvReq := updateEnvRequest{NewValue: value}
	requestJson, err := json.Marshal(updateEnvReq)

	param := map[string]string{}
	param["key"] = key

	_, err = api.SendPatch("/"+workspaceId+"/application/"+applicationId+"/env", header, param, requestJson)
	if err != nil {
		return err
	}

	return nil
}

func UpdateEnvWithLabel(workspaceId string, labels []string, key string, value string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	updateEnvReq := updateEnvRequest{NewValue: value}
	requestJson, err := json.Marshal(updateEnvReq)

	param := map[string]string{}
	param["key"] = key
	param["labels"] = strings.Join(labels, ",")

	_, err = api.SendPatch("/"+workspaceId+"/application/env", header, param, requestJson)
	if err != nil {
		return err
	}

	return nil
}
