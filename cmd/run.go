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
			return err
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

	runCmd.Flags().StringP("workspace", "w", "", "workspace id")
	runCmd.Flags().StringArrayP("label", "l", []string{}, "select labels for applications.\nwhen use this flag if you enter application id, application id will be ignored.\nif used together with the Id flag, this flag will be ignored\nex). -l test-label-1 -l test-label-2")
}
