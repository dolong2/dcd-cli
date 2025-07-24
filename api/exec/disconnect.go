package exec

import (
	"github.com/dolong2/dcd-cli/api"
)

func DisconnectDomain(workspaceId string, domainId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	_, err = api.SendPost("/"+workspaceId+"/domain/"+domainId+"/disconnect", header, map[string]string{}, []byte(""))
	if err != nil {
		return err
	}

	return nil
}
