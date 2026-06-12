package util

import (
	"encoding/json"
	"errors"
	"os"
)

func GetWorkspaceId() (string, error) {
	rawWorkspaceInfo, err := os.ReadFile("./dcd-info/workspace-info.json")
	if err != nil {
		return "", errors.New("워크스페이스 정보를 찾을수없습니다.")
	}

	var workspaceInfo map[string]interface{}

	err = json.Unmarshal(rawWorkspaceInfo, &workspaceInfo)
	if err != nil {
		return "", errors.New("워크스페이스 정보가 올바르지 않습니다.")
	}

	workspaceId := workspaceInfo["workspaceId"].(string)

	return workspaceId, nil
}