package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
)

func GetTypes() ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	result, err := api.SendGet("/application/types", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var typesResponse response.GetTypesResponse
	err = json.Unmarshal(result, &typesResponse)
	if err != nil {
		return nil, err
	}

	return typesResponse.List, nil
}
