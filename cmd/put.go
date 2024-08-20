package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// updateEnvCmd represents the put command
var updateEnvCmd = &cobra.Command{
	Use:   "update",
	Short: "update an exists env",
	Long:  `this command is used to update a exits global env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
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

// putGlobalCmd represents the put command
var putGlobalCmd = &cobra.Command{
	Use:   "put",
	Short: "update a exists global env",
	Long:  `this command is used to update a exits global env`,
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

		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both and key and value")
		}
		err := exec.UpdateGlobalEnv(workspaceId, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	envCmd.AddCommand(updateEnvCmd)

	updateEnvCmd.Flags().StringP("key", "", "", "select a key to delete")
	updateEnvCmd.Flags().StringP("workspace", "w", "", "workspace id")

	globalEnvCmd.AddCommand(putGlobalCmd)

	putGlobalCmd.Flags().StringP("key", "", "", "select a key to delete")
}
