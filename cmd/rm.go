package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm <applicationId> [flags]",
	Short: "애플리케이션의 환경변수를 삭제하는 커맨드",
	Long:  `이 커맨드는 애플리케이션의 환경변수를 삭제하는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		envKey, err := cmd.Flags().GetString("key")
		if err != nil || envKey == "" {
			return cmdError.NewCmdError(1, "환경변수 키가 입력되어야합니다.")
		}

		labels, err := cmd.Flags().GetStringArray("label")
		if err != nil {
			return err
		}

		if len(labels) != 0 {
			err := exec.RemoveEnvWithLabels(workspaceId, labels, envKey)
			if err != nil {
				return err
			}
			return nil
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		applicationId := args[0]

		err = exec.RemoveEnv(workspaceId, applicationId, envKey)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

// rmGlobalEnvCmd represents the rm command under global-env
var rmGlobalEnvCmd = &cobra.Command{
	Use:   "rm <workspaceId> [flags]",
	Short: "워크스페이스의 전역 환경변수를 삭제하는 커맨드",
	Long:  `워크스페이스의 전역 환경변수를 삭제할 수 있는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId := ""

		if len(args) == 0 {
			var err error = nil
			workspaceId, err = util.GetWorkspaceId(cmd)
			if err != nil {
				return err
			}
		} else {
			workspaceId = args[0]
		}

		envKey, err := cmd.Flags().GetString("key")
		if err != nil || envKey == "" {
			return cmdError.NewCmdError(1, "환경변수 키가 입력되어야합니다.")
		}

		err = exec.RemoveGlobalEnv(workspaceId, envKey)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	envCmd.AddCommand(rmCmd)

	rmCmd.Flags().StringP("key", "", "", "select a key to delete")
	rmCmd.Flags().StringP("workspace", "w", "", "workspace id")
	rmCmd.Flags().StringArrayP("label", "l", []string{}, "select labels for applications.\nif use this flag, you are no need to use application id.\nex). -l test-label-1 -l test-label-2")

	globalEnvCmd.AddCommand(rmGlobalEnvCmd)
	rmGlobalEnvCmd.Flags().StringP("key", "", "", "select a key to delete")
}
