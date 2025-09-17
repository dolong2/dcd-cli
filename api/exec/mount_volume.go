package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

func MountVolume(workspaceId string, volumeId string, targetApplicationId string, mountPath string, readOnly bool) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["applicationId"] = targetApplicationId

	mountVolumeRequest := request.MountVolumeRequest{
		MountPath: mountPath,
		ReadOnly:  readOnly,
	}
	requestBody, err := json.Marshal(mountVolumeRequest)
	if err != nil {
		return err
	}

	_, err = api.SendPost("/"+workspaceId+"/volume/"+volumeId+"/mount", header, param, requestBody)
	return err
}
