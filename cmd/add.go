package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <applicationId> [flags]",
	Short: "use to add an application env",
	Long:  `this command can be used to add a env to an application.`,
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
		err = exec.AddEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

// addGlobalEnvCmd represents the add command
var addGlobalEnvCmd = &cobra.Command{
	Use:   "add <workspaceId> [flags]",
	Short: "use to add a global env",
	Long:  `this command can be used to add a global env to an application.`,
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

		application := args[0]
		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both and key and value")
		}
		err := exec.AddEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	envCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("key", "k", "", "environment key")
	addCmd.Flags().StringP("value", "v", "", "environment value")
	addCmd.Flags().StringP("workspace", "w", "", "workspace id")
	globalEnvCmd.AddCommand(addGlobalEnvCmd)
	addGlobalEnvCmd.Flags().StringP("key", "k", "", "environment key")
	addGlobalEnvCmd.Flags().StringP("value", "v", "", "environment value")
}
