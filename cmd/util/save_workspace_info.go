package util

import (
	"bufio"
	"encoding/json"
	"github.com/dolong2/dcd-cli/api/exec"
	"os"
)

func SaveWorkspaceInfo(workspaceInfo exec.WorkspaceResponse) error {
	simpleWorkspaceInfo := SimpleWorkspaceInfo{WorkspaceId: workspaceInfo.Id}

	rawResult, err := json.Marshal(simpleWorkspaceInfo)
	if err != nil {
		return err
	}

	workspaceInfoDirectory := "./dcd-info"
	err = os.MkdirAll(workspaceInfoDirectory, 0755)
	if err != nil {
		return err
	}

	workspaceInfoPath := "./dcd-info/workspace-info.json"
	file, err := os.Create(workspaceInfoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(rawResult))
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
