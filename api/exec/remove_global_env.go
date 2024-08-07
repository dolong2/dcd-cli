package exec

import "github.com/dolong2/dcd-cli/api"

func RemoveGlobalEnv(workspaceId string, envKey string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	params := make(map[string]string)
	params["key"] = envKey

	_, err = api.SendDelete("/workspace/"+workspaceId+"/env", header, params)
	if err != nil {
		return err
	}

	return nil
}
