package exec

import (
	"github.com/dolong2/dcd-cli/api"
	"strings"
)

func DeleteEnv(key string, workspaceId string, applicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["key"] = key

	_, err = api.SendDelete("/"+workspaceId+"/application/"+applicationId+"/env", header, param)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEnvWithLabels(key string, workspaceId string, labels []string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["key"] = key
	param["labels"] = strings.Join(labels, ",")

	_, err = api.SendDelete("/"+workspaceId+"/application/env", header, param)
	if err != nil {
		return err
	}

	return nil
}
