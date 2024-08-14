package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

func UpdateEnv(workspaceId string, applicationId string, key string, value string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	body := map[string]string{}
	body["newValue"] = value
	envReq := envRequest{EnvList: body}
	requestJson, err := json.Marshal(envReq)

	param := map[string]string{}
	param["key"] = key

	_, err = api.SendPatch("/"+workspaceId+"/application/"+applicationId+"/env?key="+key, header, map[string]string{}, requestJson)
	if err != nil {
		return err
	}

	return nil
}
