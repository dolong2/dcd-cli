package exec

import (
	"github.com/dolong2/dcd-cli/api"
)

func DeleteGlobalEnv(key string, workspaceId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["key"] = key

	_, err = api.SendDelete("/workspace/"+workspaceId+"/env", header, param)
	if err != nil {
		return err
	}

	return nil
}
