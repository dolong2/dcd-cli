package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/dolong2/dcd-cli/api/exec/response"
	"github.com/olekukonko/tablewriter"
)

func printApplication(application response.ApplicationDetailResponse) {
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
	failureReasonDetail := []string{"Failure Reason Detail", application.FailureReasonDetail}

	table.Append(id)
	table.Append(name)
	table.Append(description)
	table.Append(applicationType)
	table.Append(githubUrl)
	for _, label := range application.Labels {
		table.Append([]string{"Label", label})
	}
	table.Append(port)
	table.Append(externalPort)
	table.Append(version)
	table.Append(status)
	table.Append(failureReason)
	table.Append(failureReasonDetail)

	table.Render()

	if len(application.InitialScripts) != 0 {
		initialScriptsTable := tablewriter.NewWriter(os.Stdout)
		initialScriptsTable.SetAutoWrapText(false)
		initialScriptsTable.SetAlignment(tablewriter.ALIGN_CENTER)

		initialScriptsTable.SetHeader([]string{"InitialScripts"})

		for _, initialScript := range application.InitialScripts {
			initialScriptsTable.Append([]string{initialScript})
		}


		initialScriptsTable.Render()	
	}
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

func printApplicationTypeList() error {
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

func printVolume(volume response.VolumeDetailResponse) {
	metaDataTable := tablewriter.NewWriter(os.Stdout)
	metaDataTable.SetAutoWrapText(false)
	metaDataTable.SetAlignment(tablewriter.ALIGN_LEFT)

	//메타데이터 출력
	metaDataTable.SetHeader([]string{"METADATA"})
	metaDataTable.Append([]string{fmt.Sprintf("ID          : %s", volume.Id)})
	metaDataTable.Append([]string{fmt.Sprintf("NAME        : %s", volume.Name)})
	metaDataTable.Append([]string{fmt.Sprintf("DESCRIPTION : %s", volume.Description)})

	metaDataTable.Render()

	if len(volume.MountList) != 0 {
		mountTable := tablewriter.NewWriter(os.Stdout)
		mountTable.SetAutoWrapText(false)
		mountTable.SetAlignment(tablewriter.ALIGN_LEFT)
		mountTable.SetAutoMergeCells(true)

		mountTable.SetHeader([]string{"MOUNT PATH", "READ ONLY", "APPLICATION ID", "APPLICATION NAME"})
		for _, mount := range volume.MountList {
			applicationInfo := mount.ApplicationInfo
			mountTable.Append([]string{mount.MountPath, strconv.FormatBool(mount.ReadOnly), applicationInfo.Id, applicationInfo.Name})
		}

		mountTable.Render()
	}

}
