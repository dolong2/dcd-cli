package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
)

type getTypesResponse struct {
	List []string `json:"list"`
}

func GetTypes() ([]string, error) {
	header := make(map[string]string)
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	header["Authorization"] = "Bearer " + accessToken

	response, err := api.SendGet("/application/types", header, map[string]string{})
	if err != nil {
		return nil, err
	}

	var typesResponse getTypesResponse
	err = json.Unmarshal(response, &typesResponse)
	if err != nil {
		return nil, err
	}

	return typesResponse.List, nil
}
