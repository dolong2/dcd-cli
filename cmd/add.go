package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "use to add an application env",
	Long:  `this command can be used to add a env to an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		application, existsAppId := cmd.Flags().GetString("application")
		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if application == "" || key == "" || value == "" || existsAppId != nil || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both application and key and value")
		}
		err := exec.AddEnv(application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	envCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("application", "", "", "input application id")
	addCmd.Flags().StringP("key", "k", "", "environment key")
	addCmd.Flags().StringP("value", "v", "", "environment value")
}
