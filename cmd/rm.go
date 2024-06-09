package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [envKey] [flags]",
	Short: "use to delete an application's env",
	Long:  `this command is used to delete an application's env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "should specify envKey")
		}
		envKey := args[0]
		applicationId, err := cmd.Flags().GetString("application")
		if err != nil || applicationId == "" {
			return cmdError.NewCmdError(1, "should specify applicationId")
		}

		err = exec.RemoveEnv(applicationId, envKey)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	envCmd.AddCommand(rmCmd)

	rmCmd.Flags().StringP("application", "", "", "select an application to delete env")
}
