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
	Short: "command to run an application",
	Long:  `this command is used to run an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		labels, err := cmd.Flags().GetStringArray("labels")
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
			return cmdError.NewCmdError(1, "must specify applicationId")
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
	runCmd.Flags().StringArrayP("labels", "l", []string(nil), "select labels for applications.\nwhen use this flag if you enter application id, application id will be ignored.\nif used together with the Id flag, this flag will be ignored\nex). -l test-label-1 -l test-label-2")
}
