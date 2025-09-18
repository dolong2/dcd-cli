package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// unmountCmd represents the unmount command
var unmountCmd = &cobra.Command{
	Use:   "unmount <volumeId>",
	Short: "볼륨 마운트를 해제하는 커맨드입니다.",
	Long:  `특정 애플리케이션의 볼륨 마운트를 해제합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		if len(args) != 1 {
			return cmdError.NewCmdError(1, "볼륨 아이디가 입력되지 않았습니다.")
		}
		volumeId := args[0]
		if volumeId == "" {
			return cmdError.NewCmdError(1, "볼륨 아이디가 입력되지 않았습니다.")
		}

		applicationId, err := cmd.Flags().GetString("application")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		if applicationId == "" {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되지 않았습니다.")
		}

		err = exec.UnmountVolume(workspaceId, volumeId, applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)

	unmountCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	unmountCmd.Flags().StringP("application", "a", "", "볼륨을 마운트할 애플리케이션 아이디")
}
