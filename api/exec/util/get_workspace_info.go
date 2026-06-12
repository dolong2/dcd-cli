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

	val, ok := workspaceInfo["workspaceId"]
    if !ok {
        return "", errors.New("워크스페이스 ID(workspaceId) 키가 존재하지 않습니다.")
    }

    workspaceId, ok := val.(string)
    if !ok {
        return "", errors.New("워크스페이스 ID가 올바른 문자열 형식이 아닙니다.")
    }

	return workspaceId, nil
}