package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify resource type")
		}
		resourceType := args[0]

		if resourceType == "applications" {
			err, done := getApplication(cmd)
			if !done {
				return err
			}
		} else if resourceType == "workspaces" {
			err := getWorkspace(cmd)
			if err != nil {
				return err
			}
		}

		return nil
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
	workspaceId, err := util.GetWorkspaceId(cmd)
	if err != nil {
		return err, false
	}

	applicationId, err := cmd.Flags().GetString("id")
	if applicationId == "" || err != nil {
		applications, err := exec.GetApplications(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error()), false
		}

		printApplicationList(applications.Applications)

		return nil, true
	}

	application, err := exec.GetApplication(workspaceId, applicationId)
	if err != nil {
		return cmdError.NewCmdError(1, err.Error()), false
	}

	printApplication(*application)

	return nil, true
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("workspace", "w", "", "used to get resources in a workspace")
	getCmd.Flags().StringP("id", "", "", "specify resource id")
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

func printApplicationList(applicationList []exec.ApplicationResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Description", "Application Type", "Github URL", "Port", "External Port", "Version", "Status"})

	for _, application := range applicationList {
		row := []string{application.Id, application.Name, application.Description, application.ApplicationType, application.GithubUrl, strconv.Itoa(application.Port), strconv.Itoa(application.ExternalPort), application.Version, application.Status}
		table.Append(row)
	}

	table.Render()
}
