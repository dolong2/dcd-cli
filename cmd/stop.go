package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop <applicationId> [flags]",
	Short: "애플리케이션의 동작을 정지하는 커맨드",
	Long:  `애플리케이션이 현재 동작한다면 애플리케이션의 동작을 정지시키는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		labels, err := cmd.Flags().GetStringArray("label")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		if len(labels) != 0 {
			err := exec.StopApplicationWithLabels(workspaceId, labels)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			return nil
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		applicationId := args[0]
		err = exec.StopApplication(workspaceId, applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	stopCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	stopCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n이 플래그를 사용한다면 명시한 애플리케이션 아이디는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
}
