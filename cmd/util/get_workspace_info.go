package util

import (
	"encoding/json"
	"errors"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"os"
)

func GetWorkspaceId(cmd *cobra.Command) (string, error) {
	workspaceId, err := cmd.Flags().GetString("workspace")
	if err != nil || workspaceId == "" {
		workspaceId, err = getWorkspaceInfo()
		if err != nil {
			return "", cmdError.NewCmdError(1, "워크스페이스 아이디가 입력되어야합니다.")
		}
	}
	return workspaceId, nil
}

func getWorkspaceInfo() (string, error) {
	rawWorkspaceInfo, err := os.ReadFile("./dcd-info/workspace-info.json")
	if err != nil {
		return "", errors.New("워크스페이스 정보를 찾을 수 없습니다.")
	}

	var simpleWorkspaceInfo SimpleWorkspaceInfo

	err = json.Unmarshal(rawWorkspaceInfo, &simpleWorkspaceInfo)
	if err != nil {
		return "", errors.New("옳바르지 않은 워크스페이스 정보입니다.")
	}

	return simpleWorkspaceInfo.WorkspaceId, nil
}
