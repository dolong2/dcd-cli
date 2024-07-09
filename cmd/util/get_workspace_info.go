package util

import (
	"encoding/json"
	"errors"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"os"
)

func GetWorkspaceId(cmd *cobra.Command) (string, error) {
	workspaceId, err := getWorkspaceInfo()
	if err != nil {
		workspaceFlag, err := cmd.Flags().GetString("workspace")
		if workspaceFlag != "" || err != nil {
			return "", cmdError.NewCmdError(1, "must specify workspace id")
		}
		workspaceId = workspaceFlag
	}
	return workspaceId, nil
}

func getWorkspaceInfo() (string, error) {
	rawWorkspaceInfo, err := os.ReadFile("./dcd-info/workspace-info.json")
	if err != nil {
		return "", errors.New("not found workspace info")
	}

	var simpleWorkspaceInfo SimpleWorkspaceInfo

	err = json.Unmarshal(rawWorkspaceInfo, &simpleWorkspaceInfo)
	if err != nil {
		return "", errors.New("invalid workspace info")
	}

	return simpleWorkspaceInfo.WorkspaceId, nil
}
