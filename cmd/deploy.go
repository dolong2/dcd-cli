package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy <applicationId>",
	Short: "command to deploy an application",
	Long:  `this command is used to deploy an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		applicationId := args[0]
		err := exec.DeployApplication(applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
