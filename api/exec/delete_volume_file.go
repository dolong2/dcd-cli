package exec

import "github.com/dolong2/dcd-cli/api"

func DeleteVolumeFile(workspaceId string, volumeId string, targetPath string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	param := make(map[string]string)
	param["path"] = targetPath

	_, err = api.SendDelete("/"+workspaceId+"/volume/"+volumeId+"/files", header, param)
	return err
}
