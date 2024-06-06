package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"

	"github.com/spf13/cobra"
)

// applicationsCmd represents the applications command
var applicationsCmd = &cobra.Command{
	Use:   "applications",
	Short: "sub command to get applications",
	Long:  `this command can be used to get workspaces`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, existsWorkspaceId := cmd.Flags().GetString("workspace")
		if existsWorkspaceId != nil || workspaceId == "" {
			return cmdError.NewCmdError(1, "requires workspace flag to get your applications")
		}
		applicationListResponse, err := exec.GetApplications(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		for _, application := range applicationListResponse.Applications {
			printApplication(application)
		}
		return nil
	},
}

func init() {
	getCmd.AddCommand(applicationsCmd)
	applicationsCmd.Flags().StringP("workspace", "w", "", "use to identify workspace")
	applicationsCmd.Flags().StringP("id", "", "", "use to get an application by applicationId")
}

func printApplication(application exec.ApplicationResponse) {
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
	fmt.Println()
	fmt.Println()
}
