package exec

import "github.com/dolong2/dcd-cli/api"

func Logout() error {
	header := make(map[string]string)
	token, err := GetAccessToken()
	if err != nil {
		return err
	}
	header["Authorization"] = "Bearer " + token

	_, err = api.SendDelete("/auth", header, map[string]string{})
	if err != nil {
		return err
	}

	return nil
}
