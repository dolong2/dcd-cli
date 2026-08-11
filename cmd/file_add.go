package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// fileAddCmd represents the file add command
var fileAddCmd = &cobra.Command{
	Use:   "add <volumeId> <filePath>",
	Short: "볼륨에 파일을 업로드하는 커맨드입니다.",
	Long:  `로컬 파일을 볼륨의 특정 경로에 업로드합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return cmdError.NewCmdError(1, "볼륨 아이디 혹은 업로드할 파일 경로가 입력되지 않았습니다.")
		}
		volumeId := args[0]
		localFilePath := args[1]

		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		targetPath, err := cmd.Flags().GetString("path")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		if targetPath == "" {
			return cmdError.NewCmdError(1, "볼륨 내 저장할 경로가 입력되지 않았습니다.")
		}

		createDirectory, err := cmd.Flags().GetBool("createDirectory")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		err = exec.UploadVolumeFile(workspaceId, volumeId, localFilePath, targetPath, createDirectory)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	fileCmd.AddCommand(fileAddCmd)

	fileAddCmd.Flags().StringP("path", "p", "", "볼륨 내 저장할 파일 경로")
	fileAddCmd.Flags().BoolP("directory", "d", false, "상위 디렉토리가 없을 경우 생성 여부")
}
