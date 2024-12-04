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
	Short: "리소스를 조회하는 커맨드",
	Long: `이 커맨드는 리소스를 조회하기 위해서 사용됩니다.
리소스 타입:
	workspaces - 이 리소스 타입은 여러 애플리케이션을 가지고, 작업구역(네트워크)를 나눌때 사용합니다.
	applications - 이 리소스 타입은 특정 라이브러리 혹은 프레임워크가 컨테이너에서 동작하게 하는 리소스 타입입니다.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "리소스 타입이 입력되어야 합니다.")
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
		} else {
			return cmdError.NewCmdError(1, "올바르지 않은 리소스 타입입니다.")
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
		labels, err := cmd.Flags().GetStringArray("label")

		var applications *exec.ApplicationListResponse

		// id, labels 플래그 둘다 없을때, 조건 없이 애플리케이션 조회
		if len(labels) == 0 || err != nil {
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

	getCmd.Flags().StringP("workspace", "w", "", "리소스를 가져올 워크스페이스 아이디")
	getCmd.Flags().StringP("id", "", "", "리소스 아이디")
	getCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n워크스페이스를 가져올때 해당 플래그를 사용하면, 해당 플래그는 무시됩니다.\n리소스 아이디를 사용한다면, 이 커맨드는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
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
