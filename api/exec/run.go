package exec

import "github.com/dolong2/dcd-cli/api"

func RunApplication(applicationId string) error {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + accessToken

	_, err = api.SendPost("/application/"+applicationId+"/run", header, []byte(""))
	return err
}
