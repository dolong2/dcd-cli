package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resourceType> [flags]",
	Short: "command to get resources",
	Long: `this command is used to get resources.
resourceType:
	workspaces - this resource type refers a resource that has applications and be sectioned networks
	applications - this resource type refers a resource that works something in a container
`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
