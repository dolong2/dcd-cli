package cmd

import (
	"github.com/spf13/cobra"
)

// globalEnvCmd represents the globalEnv command
var globalEnvCmd = &cobra.Command{
	Use:   "global-env",
	Short: "manage workspace global environment variables",
	Long:  `command to manage global env`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(globalEnvCmd)
}
