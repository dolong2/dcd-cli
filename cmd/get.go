package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/dolong2/dcd-cli/api/exec/response"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/resource"
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
조회 가능한 리소스 타입:
	workspaces - 이 리소스 타입은 여러 애플리케이션을 가지고, 작업구역을 나눌때 사용합니다.
	applications - 이 리소스 타입은 특정 라이브러리 혹은 프레임워크가 컨테이너에서 동작하게 하는 리소스 타입입니다.
	types - 애플리케이션의 타입 종류를 나타내는 리소스 타입입니다.
	domains - 해당 리소스타입은 애플리케이션을 HTTPS로 외부에 공개할때 사용되는 리소스 타입입니다.
	envs - 애플리케이션에서 사용될 수 있는 환경변수를 나타내는 리소스 타입입니다.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "리소스 타입이 입력되어야 합니다.")
		}
		resourceType := resource.Type(args[0])
		if !resourceType.IsValid() {
			return cmdError.NewCmdError(1, "올바르지 않은 리소스 타입입니다.")
		}

		switch {
		case resourceType.IsEqual(resource.Application):
			err := getApplication(cmd)
			if err != nil {
				return err
			}
		case resourceType.IsEqual(resource.Workspace):
			err := getWorkspace(cmd)
			if err != nil {
				return err
			}
		case resourceType.IsEqual(resource.ApplicationType):
			err := printApplicationTypes()
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		case resourceType.IsEqual(resource.Domain):
			err := getDomain(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		case resourceType.IsEqual(resource.Env):
			err := getEnv(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		default:
			return cmdError.NewCmdError(1, "조회할 수 없는 리소스 타입입니다.")
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

		var applications *response.ApplicationListResponse

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

func getDomain(cmd *cobra.Command) error {
	workspaceId, err := util.GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	domainListResponse, err := exec.GetDomains(workspaceId)
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	}

	printDomainList(domainListResponse.Domains)

	return nil
}

func getEnv(cmd *cobra.Command) error {
	workspaceId, err := util.GetWorkspaceId(cmd)
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	}

	envId, err := cmd.Flags().GetString("id")
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	} else if envId == "" {
		envListResponse, err := exec.GetEnvList(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		printEnvList(*envListResponse)
	} else {
		envResponse, err := exec.GetEnv(workspaceId, envId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		printEnv(*envResponse)
	}

	return nil
}

func getVolume(cmd *cobra.Command) error {
	workspaceId, err := util.GetWorkspaceId(cmd)
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	}

	volumeId, err := cmd.Flags().GetString("id")
	if err != nil {
		return cmdError.NewCmdError(1, err.Error())
	} else if volumeId == "" {
		volumeList, err := exec.GetVolumeList(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		printVolumeList(*volumeList)
	} else {

	}

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("workspace", "w", "", "리소스를 가져올 워크스페이스 아이디")
	getCmd.Flags().StringP("id", "", "", "리소스 아이디")
	getCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n워크스페이스를 가져올때 해당 플래그를 사용하면, 해당 플래그는 무시됩니다.\n리소스 아이디를 사용한다면, 이 플래그는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
}

func printApplication(application response.ApplicationResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	id := []string{"ID", application.Id}
	name := []string{"Name", application.Name}
	description := []string{"Description", application.Description}
	applicationType := []string{"Application Type", application.ApplicationType}
	githubUrl := []string{"GitHub Url", application.GithubUrl}
	port := []string{"Port", strconv.Itoa(application.Port)}
	externalPort := []string{"External Port", strconv.Itoa(application.ExternalPort)}
	version := []string{"Version", application.Version}
	status := []string{"Status", application.Status}
	failureReason := []string{"Failure Reason", application.FailureReason}

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
	table.Append(failureReason)

	table.Render()
}

func printApplicationList(applicationList []response.ApplicationResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	table.SetHeader([]string{"ID", "Name", "Description", "Application Type", "Github URL", "Port", "External Port", "Version", "Status", "Labels"})

	for _, application := range applicationList {
		labels := application.Labels

		labelStr := ""
		if len(labels) > 2 {
			labelStr = strings.Join(labels[:2], ", ")
			labelStr += " ..."
		} else {
			labelStr = strings.Join(labels, ", ")
		}

		row := []string{
			application.Id,
			application.Name,
			application.Description,
			application.ApplicationType,
			application.GithubUrl,
			strconv.Itoa(application.Port),
			strconv.Itoa(application.ExternalPort),
			application.Version,
			application.Status,
			labelStr,
		}
		table.Append(row)
	}

	table.Render()
}

func printWorkspace(workspace response.WorkspaceDetailResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	id := []string{"ID", workspace.Id}
	title := []string{"Name", workspace.Title}
	description := []string{"Description", workspace.Description}

	table.Append(id)
	table.Append(title)
	table.Append(description)

	table.Render()
}

func printWorkspaceList(workspaceList []response.WorkspaceResponse, usedWorkspaceId string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

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

func printApplicationTypes() error {
	types, err := exec.GetTypes()

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Types"})

	for _, typeValue := range types {
		table.Append([]string{typeValue})
	}

	table.Render()

	return nil
}

func printDomainList(domainList []response.DomainResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	table.SetHeader([]string{"ID", "NAME", "Description", "STATUS", "APPLICATION"})

	for _, domain := range domainList {

		var status, applicationName string
		if domain.Application != nil {
			status = "CONNECTED"
			applicationName = domain.Application.Name
		} else {
			status = "UNCONNECTED"
			applicationName = ""
		}

		row := []string{domain.DomainId, domain.Name, domain.Description, status, applicationName}
		table.Append(row)
	}

	table.Render()
}

func printEnv(env response.EnvResponse) {
	metaDataTable := tablewriter.NewWriter(os.Stdout)
	metaDataTable.SetAutoWrapText(false)
	metaDataTable.SetAlignment(tablewriter.ALIGN_LEFT)

	//메타데이터 출력
	metaDataTable.SetHeader([]string{"METADATA"})
	metaDataTable.Append([]string{fmt.Sprintf("ID          : %s", env.Id)})
	metaDataTable.Append([]string{fmt.Sprintf("NAME        : %s", env.Name)})
	metaDataTable.Append([]string{fmt.Sprintf("DESCRIPTION : %s", env.Description)})

	metaDataTable.Render()

	detailTable := tablewriter.NewWriter(os.Stdout)
	detailTable.SetAutoWrapText(false)
	detailTable.SetAlignment(tablewriter.ALIGN_LEFT)

	//환경변수 디테일 출력
	detailTable.SetHeader([]string{"KEY", "VALUE", "ENCRYPTION"})
	for _, detail := range env.Details {
		detailTable.Append([]string{detail.Key, detail.Value, fmt.Sprintf("%v", detail.Encryption)})
	}

	detailTable.Render()
}

func printEnvList(envList response.EnvListResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader([]string{"ID", "NAME", "DESCRIPTION"})
	for _, envSimpleResponse := range envList.List {
		table.Append([]string{envSimpleResponse.Id, envSimpleResponse.Name, envSimpleResponse.Description})
	}

	table.Render()
}

func printVolumeList(volumeList response.VolumeListResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	table.SetHeader([]string{"ID", "Name", "Description"})

	for _, volume := range volumeList.List {
		row := []string{volume.Id, volume.Name, volume.Description}
		table.Append(row)
	}

	table.Render()
}
