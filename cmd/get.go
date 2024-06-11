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

func getWorkspace(cmd *cobra.Command) error {
	id, existsFlagErr := cmd.Flags().GetString("id")
	if id == "" || existsFlagErr != nil {
		workspaceList, err := exec.GetWorkspaces()
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		for _, workspace := range workspaceList.List {
			fmt.Printf("ID: %s\nTitle: %s\nDescription: %s\n\n", workspace.Id, workspace.Title, workspace.Description)
		}
	} else {
		workspace, err := exec.GetWorkspace(id)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		fmt.Printf("ID: %s\nTitle: %s\nDescription: %s\n\n", workspace.Id, workspace.Title, workspace.Description)
	}
	return nil
}
}

func init() {
	rootCmd.AddCommand(getCmd)
}
