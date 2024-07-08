package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "command to get application logs",
	Long:  `this command is used to get application logs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId()
		if err != nil {
			workspaceFlag, err := cmd.Flags().GetString("workspace")
			if workspaceFlag != "" || err != nil {
				return cmdError.NewCmdError(1, "must specify workspace id")
			}
			workspaceId = workspaceFlag
		}

		if len(args) == 0 || args[0] == "" {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		applicationId := args[0]
		logs, err := exec.GetLog(workspaceId, applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		for _, log := range logs {
			fmt.Println(log)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	logsCmd.Flags().StringP("workspace", "w", "", "workspace id")
}
