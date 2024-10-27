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
	Short: "command to stop an application",
	Long:  `this command is used to stop an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
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

	stopCmd.Flags().StringP("workspace", "w", "", "workspace id")
	stopCmd.Flags().StringArrayP("label", "l", []string{}, "select labels for applications.\nwhen use this flag if you enter application id, application id will be ignored.\nif used together with the Id flag, this flag will be ignored\nex). -l test-label-1 -l test-label-2")
}
