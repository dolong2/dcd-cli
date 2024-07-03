package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <applicationId> [flags]",
	Short: "use to update an application env",
	Long:  `this command can be used to update a env to an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId()
		if err != nil {
			workspaceFlag, err := cmd.Flags().GetString("workspace")
			if workspaceFlag != "" || err != nil {
				return cmdError.NewCmdError(1, "must specify workspace id")
			}
			workspaceId = workspaceFlag
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		application := args[0]
		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both and key and value")
		}
		err = exec.UpdateEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	envCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("key", "k", "", "environment key")
	updateCmd.Flags().StringP("value", "v", "", "environment value")
	updateCmd.Flags().StringP("workspace", "w", "", "workspace id")
}
