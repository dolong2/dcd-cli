package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

func UpdateGlobalEnv(workspaceId string, key string, value string) error {
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

	_, err = api.SendPatch("/workspace/"+workspaceId+"/env?key="+key, header, map[string]string{}, requestJson)
	if err != nil {
		return err
	}

	return nil
}
