package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

func ConnectDomain(workspaceId string, domainId string, applicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	connectDomainRequest := request.ConnectDomainRequest{ApplicationId: applicationId}
	jsonConnectDomainRequest, err := json.Marshal(connectDomainRequest)
	if err != nil {
		return err
	}

	_, err = api.SendPost("/"+workspaceId+"/domain/"+domainId+"/connect", header, map[string]string{}, jsonConnectDomainRequest)
	if err != nil {
		return err
	}

	return nil
}
