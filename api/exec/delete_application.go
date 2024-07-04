package exec

import "github.com/dolong2/dcd-cli/api"

func DeleteApplication(workspaceId string, applicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	_, err = api.SendDelete("/"+workspaceId+"/application/"+applicationId, header, map[string]string{})
	return err
}
