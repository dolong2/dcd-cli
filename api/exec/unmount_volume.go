package exec

import (
	"github.com/dolong2/dcd-cli/api"
)

func UnmountVolume(workspaceId string, volumeId string, targetApplicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["applicationId"] = targetApplicationId

	_, err = api.SendDelete("/"+workspaceId+"/volume/"+volumeId+"/mount", header, param)
	return err
}
