package util

import (
	"encoding/json"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"os"
)

func GetWorkspaceId() (string, error) {
	rawWorkspaceInfo, err := os.ReadFile("./dcd-info/workspace-info.json")
	if err != nil {
		return "", cmdError.NewCmdError(125, "not found workspace info")
	}

	var simpleWorkspaceInfo SimpleWorkspaceInfo

	err = json.Unmarshal(rawWorkspaceInfo, &simpleWorkspaceInfo)
	if err != nil {
		return "", cmdError.NewCmdError(125, "invalid workspace info")
	}

	return simpleWorkspaceInfo.WorkspaceId, nil
}
