package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <applicationId> [flags]",
	Short: "애플리케이션을 실행하는 커맨드",
	Long:  `애플리케이션을 실행시키는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		labels, err := cmd.Flags().GetStringArray("label")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		if len(labels) != 0 {
			err := exec.RunApplicationWithLabels(workspaceId, labels)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			return nil
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		applicationId := args[0]
		err = exec.RunApplication(workspaceId, applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	runCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n이 플래그를 사용할때 명시한 애플리케이션 아이디는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
}
