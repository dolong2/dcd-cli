package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// fileRemoveCmd represents the file remove command
var fileRemoveCmd = &cobra.Command{
	Use:   "remove <volumeId>",
	Short: "볼륨에서 파일을 제거하는 커맨드입니다.",
	Long:  `볼륨의 특정 경로에 있는 파일을 제거합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmdError.NewCmdError(1, "볼륨 아이디가 입력되지 않았습니다.")
		}
		volumeId := args[0]

		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		targetPath, err := cmd.Flags().GetString("path")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		if targetPath == "" {
			return cmdError.NewCmdError(1, "볼륨 내 제거할 파일 경로가 입력되지 않았습니다.")
		}

		err = exec.DeleteVolumeFile(workspaceId, volumeId, targetPath)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	fileCmd.AddCommand(fileRemoveCmd)

	fileRemoveCmd.Flags().StringP("path", "p", "", "볼륨에서 제거할 파일 경로")
}
