package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

func SetDomain(workspaceId string, applicationId string, domain string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	domainRequest := request.SetDomainRequest{Domain: domain}
	requestBody, err := json.Marshal(domainRequest)
	if err != nil {
		return err
	}

	_, err = api.SendPost("/"+workspaceId+"/application/"+applicationId+"/domain", header, map[string]string{}, requestBody)
	if err != nil {
		return err
	}

	return nil
}
