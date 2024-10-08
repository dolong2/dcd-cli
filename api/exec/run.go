package exec

import (
	"github.com/dolong2/dcd-cli/api"
)

func RunApplication(workspaceId string, applicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	_, err = api.SendPost("/"+workspaceId+"/application/"+applicationId+"/run", header, map[string]string{}, []byte(""))
	return err
}
