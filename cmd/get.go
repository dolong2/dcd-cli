package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
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

func getApplication(cmd *cobra.Command) (error, bool) {
	workspaceId, existsWorkspaceId := cmd.Flags().GetString("workspace")
	applicationId, existsApplicationId := cmd.Flags().GetString("id")
	if (existsWorkspaceId != nil || workspaceId == "") && (existsApplicationId != nil || applicationId == "") {
		return cmdError.NewCmdError(1, "requires workspace flag to get your applications"), true
	}
	if workspaceId != "" && applicationId != "" {
		return cmdError.NewCmdError(125, "simultaneous use of workspaceId and application is not permitted"), true
	}
	if workspaceId != "" {
		applicationListResponse, err := exec.GetApplications(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error()), true
		}
		for _, application := range applicationListResponse.Applications {
			printApplication(application)
		}
	} else if applicationId != "" {
		applicationResponse, err := exec.GetApplication(applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error()), true
		}
		printApplication(*applicationResponse)
	}
	return nil, false
}

func init() {
	rootCmd.AddCommand(getCmd)
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
