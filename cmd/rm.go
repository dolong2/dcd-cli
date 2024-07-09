package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm <applicationId> [flags]",
	Short: "use to delete an application's env",
	Long:  `this command is used to delete an application's env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		applicationId := args[0]
		envKey, err := cmd.Flags().GetString("key")
		if err != nil || envKey == "" {
			return cmdError.NewCmdError(1, "should specify envKey")
		}

		err = exec.RemoveEnv(workspaceId, applicationId, envKey)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	envCmd.AddCommand(rmCmd)

	rmCmd.Flags().StringP("key", "", "", "select an key to delete")
	rmCmd.Flags().StringP("workspace", "w", "", "workspace id")
}
