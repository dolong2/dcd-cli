package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy <applicationId>",
	Short: "command to deploy an application",
	Long:  `this command is used to deploy an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("deploy called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
