package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/spf13/cobra"
)

// workspacesCmd represents the workspace command
var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "sub command to get workspaces",
	Long:  `this command can be used to get workspaces`,
	Run: func(cmd *cobra.Command, args []string) {
		if err != nil {
			fmt.Println(err)
		}
		for _, workspace := range workspaceList.List {
			workspaceList, err := exec.GetWorkspaces()
			fmt.Printf("ID: %s\nTitle: %s\nDescription: %s\n\n", workspace.Id, workspace.Title, workspace.Description)
		}
	},
}

func init() {
	getCmd.AddCommand(workspacesCmd)
	workspacesCmd.Flags().StringP("id", "", "", "use this flag to get a workspace by ID")
}
