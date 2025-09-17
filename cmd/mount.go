package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount <volumeId> <mountPath>",
	Short: "볼륨을 애플리케이션에 마운트할때 사용하는 커맨드입니다.",
	Long:  `존재하는 볼륨을 애플리케이션에 마운트합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		if len(args) != 2 {
			return cmdError.NewCmdError(1, "볼륨 아이디 혹은 마운트 경로가 입력되지 않았습니다.")
		}
		volumeId := args[0]
		if volumeId == "" {
			return cmdError.NewCmdError(1, "볼륨 아이디가 입력되지 않았습니다.")
		}
		mountPath := args[1]
		if mountPath == "" {
			return cmdError.NewCmdError(1, "마운트 경로가 입력되지 않았습니다.")
		}

		applicationId, err := cmd.Flags().GetString("application")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		if applicationId == "" {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되지 않았습니다.")
		}

		readOnly, err := cmd.Flags().GetBool("readOnly")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		err = exec.MountVolume(workspaceId, volumeId, applicationId, mountPath, readOnly)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)

	mountCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	mountCmd.Flags().StringP("application", "a", "", "볼륨을 마운트할 애플리케이션 아이디")
	mountCmd.Flags().BoolP("readOnly", "", false, "읽기 전용으로 마운트")
}
