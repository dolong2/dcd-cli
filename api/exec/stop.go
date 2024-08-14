package exec

import "github.com/dolong2/dcd-cli/api"

func StopApplication(workspaceId string, applicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	_, err = api.SendPost("/"+workspaceId+"/application/"+applicationId+"/stop", header, map[string]string{}, []byte(""))
	return err
}
