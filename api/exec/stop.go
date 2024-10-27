package exec

import (
	"github.com/dolong2/dcd-cli/api"
	"strings"
)

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

func StopApplicationWithLabels(workspaceId string, labels []string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	params := make(map[string]string)
	params["labels"] = strings.Join(labels, ",")
	_, err = api.SendPost("/"+workspaceId+"/application/stop", header, params, []byte(""))
	return err
}
