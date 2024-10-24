package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
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
			err := getApplication(cmd)
			if err != nil {
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

		usedWorkspaceId, usedWorkspace := util.GetWorkspaceId(cmd)
		if usedWorkspace != nil {
			usedWorkspaceId = ""
		}

		printWorkspaceList(workspaceList.List, usedWorkspaceId)

		return nil
	}

	workspace, err := exec.GetWorkspace(id)
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	}

	printWorkspace(*workspace)

	return nil
}

func getApplication(cmd *cobra.Command) error {
	workspaceId, err := util.GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	applicationId, err := cmd.Flags().GetString("id")

	if applicationId == "" && err == nil {
		labels, err := cmd.Flags().GetString("labels")

		var applications *exec.ApplicationListResponse

		// id, labels 플래그 둘다 없을때, 조건 없이 애플리케이션 조회
		if labels == "" || err != nil {
			applications, err = exec.GetApplications(workspaceId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			printApplicationList(applications.Applications)

			return nil
		} else {
			applications, err = exec.GetApplicationsByLabels(workspaceId, labels)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		}

		printApplicationList(applications.Applications)

		return nil
	}

	application, err := exec.GetApplication(workspaceId, applicationId)
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	}

	printApplication(*application)

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("workspace", "w", "", "used to get resources in a workspace")
	getCmd.Flags().StringP("id", "", "", "specify resource id")
	getCmd.Flags().StringP("labels", "l", "", "select labels for applications.\nif use this flag when get workspaces, this flag will be ignored.\nif used together with the Id flag, this flag will be ignored\nex). test-label-1,test-label-2")
}

func printApplication(application exec.ApplicationResponse) {
	table := tablewriter.NewWriter(os.Stdout)

	id := []string{"ID", application.Id}
	name := []string{"Name", application.Name}
	description := []string{"Description", application.Description}
	applicationType := []string{"Application Type", application.ApplicationType}
	githubUrl := []string{"GitHub Url", application.GithubUrl}
	port := []string{"Port", strconv.Itoa(application.Port)}
	externalPort := []string{"External Port", strconv.Itoa(application.ExternalPort)}
	version := []string{"Version", application.Version}
	status := []string{"Status", application.Status}

	table.Append(id)
	table.Append(name)
	table.Append(description)
	table.Append(applicationType)
	table.Append(githubUrl)
	for key, value := range application.Env {
		env := []string{"ENV", key + " : " + value}
		table.Append(env)
	}
	for _, label := range application.Labels {
		table.Append([]string{"Label", label})
	}
	table.Append(port)
	table.Append(externalPort)
	table.Append(version)
	table.Append(status)

	table.Render()
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

func printWorkspace(workspace exec.WorkspaceDetailResponse) {
	table := tablewriter.NewWriter(os.Stdout)

	id := []string{"ID", workspace.Id}
	title := []string{"Name", workspace.Title}
	description := []string{"Description", workspace.Description}
	envString := ""
	for key, value := range workspace.GlobalEnv {
		envString += "{ " + key + " : " + value + " } , "
	}
	envString = envString[:len(envString)-1]
	globalEnv := []string{"Global Env", strings.Trim(envString, ",")}

	table.Append(id)
	table.Append(title)
	table.Append(description)
	table.Append(globalEnv)

	table.Render()
}

func printWorkspaceList(workspaceList []exec.WorkspaceResponse, usedWorkspaceId string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{" ", "ID", "TITLE", "Description"})

	for _, workspace := range workspaceList {
		usedWorkspace := ""
		if usedWorkspaceId == workspace.Id {
			usedWorkspace = "*"
		}

		row := []string{usedWorkspace, workspace.Id, workspace.Title, workspace.Description}
		table.Append(row)
	}

	table.Render()
}
