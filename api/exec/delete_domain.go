package exec

import (
	"github.com/dolong2/dcd-cli/api"
)

func DeleteDomain(workspaceId string, domainId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	_, err = api.SendDelete("/"+workspaceId+"/domain/"+domainId, header, map[string]string{})
	if err != nil {
		return err
	}
	return nil
}
