package exec

import (
	"github.com/dolong2/dcd-cli/api"
	"strings"
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

func RunApplicationWithLabels(workspaceId string, labels []string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	labelValues := strings.Join(labels, ",")
	params := make(map[string]string)
	params["labels"] = labelValues

	_, err = api.SendPost("/"+workspaceId+"/application/run", header, params, []byte(""))
	return err
}
