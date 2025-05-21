package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetProfile() (*response.Profile, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/user/profile", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	profile := response.Profile{}
	err = json.Unmarshal(result, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
