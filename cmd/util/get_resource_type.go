package util

import (
	"errors"

	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/dolong2/dcd-cli/api/exec/response"
	"github.com/dolong2/dcd-cli/cmd/resource"
	"github.com/spf13/cobra"
)

func GetResource(cmd *cobra.Command, resourceType resource.Type) (error) {
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
			err := getApplicationType()
			if err != nil {
				return err
			}
		case resourceType.IsEqual(resource.Domain):
			err := getDomain(cmd)
			if err != nil {
				return err
			}
		case resourceType.IsEqual(resource.Env):
			err := getEnv(cmd)
			if err != nil {
				return err
			}
		case resourceType.IsEqual(resource.VOLUME):
			err := getVolume(cmd)
			if err != nil {
				return err
			}
		default:
			return errors.New("조회할 수 없는 리소스 타입입니다.")
		}
	return nil
}

func getWorkspace(cmd *cobra.Command) error {
	id, existsFlagErr := cmd.Flags().GetString("id")
	if id == "" || existsFlagErr != nil {
		workspaceList, err := exec.GetWorkspaces()
		if err != nil {
			return err
		}

		usedWorkspaceId, usedWorkspace := GetWorkspaceId(cmd)
		if usedWorkspace != nil {
			usedWorkspaceId = ""
		}

		printWorkspaceList(workspaceList.List, usedWorkspaceId)

		return nil
	}

	workspace, err := exec.GetWorkspace(id)
	if err != nil {
		return err
	}

	printWorkspace(*workspace)

	return nil
}

func getApplication(cmd *cobra.Command) error {
	workspaceId, err := GetWorkspaceId(cmd)
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
				return err
			}

			printApplicationList(applications.Applications)

			return nil
		} else {
			applications, err = exec.GetApplicationsByLabels(workspaceId, labels)
			if err != nil {
				return err
			}
		}

		printApplicationList(applications.Applications)

		return nil
	}

	application, err := exec.GetApplication(workspaceId, applicationId)
	if err != nil {
		return err
	}

	printApplication(*application)

	return nil
}

func getApplicationType() error {
	err := printApplicationTypeList()

	if err != nil {
		return err
	}

	return nil
}

func getDomain(cmd *cobra.Command) error {
	workspaceId, err := GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	domainListResponse, err := exec.GetDomains(workspaceId)
	if err != nil {
		return err
	}

	printDomainList(domainListResponse.Domains)

	return nil
}

func getEnv(cmd *cobra.Command) error {
	workspaceId, err := GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	envId, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
	} else if envId == "" {
		envListResponse, err := exec.GetEnvList(workspaceId)
		if err != nil {
			return err
		}

		printEnvList(*envListResponse)
	} else {
		envResponse, err := exec.GetEnv(workspaceId, envId)
		if err != nil {
			return err
		}

		printEnv(*envResponse)
	}

	return nil
}

func getVolume(cmd *cobra.Command) error {
	workspaceId, err := GetWorkspaceId(cmd)
	if err != nil {
		return err
	}

	volumeId, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
	} else if volumeId == "" {
		volumeList, err := exec.GetVolumeList(workspaceId)
		if err != nil {
			return err
		}
		printVolumeList(*volumeList)
	} else {
		volumeDetail, err := exec.GetVolume(workspaceId, volumeId)
		if err != nil {
			return err
		}
		printVolume(*volumeDetail)
	}

	return nil
}