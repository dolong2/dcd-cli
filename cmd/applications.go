package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"

	"github.com/spf13/cobra"
)

// applicationsCmd represents the applications command
var applicationsCmd = &cobra.Command{
	Use:   "applications",
	Short: "sub command to get applications",
	Long:  `this command can be used to get workspaces`,
	Run: func(cmd *cobra.Command, args []string) {
		workspaceId, err2 := cmd.Flags().GetString("workspace")
		if err2 != nil {
			fmt.Println(err2)
		}
		applicationListResponse, err := exec.GetApplications(workspaceId)
		if err != nil {
			fmt.Println(err)
		}
		for _, application := range applicationListResponse.Applications {
			fmt.Printf("ID: %s\n", application.Id)
			fmt.Printf("Name: %s\n", application.Name)
			fmt.Printf("Description: %s\n", application.Description)
			fmt.Printf("Application Type: %s\n", application.ApplicationType)
			fmt.Printf("GitHub URL: %s\n", application.GithubUrl)
			fmt.Printf("Environment Variables: [\n")
			for key, value := range application.Env {
				fmt.Printf("  %s: %s\n", key, value)
			}
			fmt.Printf("]\n")
			fmt.Printf("Port: %d\n", application.Port)
			fmt.Printf("External Port: %d\n", application.ExternalPort)
			fmt.Printf("Version: %s\n", application.Version)
			fmt.Printf("Status: %s\n", application.Status)
		}
	},
}

func init() {
	getCmd.AddCommand(applicationsCmd)
	applicationsCmd.Flags().StringP("workspace", "w", "", "use to identify workspace")
}
