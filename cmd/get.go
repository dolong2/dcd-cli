package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/dolong2/dcd-cli/api/exec/response"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/resource"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resourceType> [flags]",
	Short: "리소스를 조회하는 커맨드",
	Long: `이 커맨드는 리소스를 조회하기 위해서 사용됩니다.
조회 가능한 리소스 타입:
	workspace(ws) - 이 리소스 타입은 여러 애플리케이션을 가지고, 작업구역을 나눌때 사용합니다.
	application(app) - 이 리소스 타입은 특정 라이브러리 혹은 프레임워크가 컨테이너에서 동작하게 하는 리소스 타입입니다.
	type(ts) - 애플리케이션의 타입 종류를 나타내는 리소스 타입입니다.
	domain(dom) - 해당 리소스타입은 애플리케이션을 HTTPS로 외부에 공개할때 사용되는 리소스 타입입니다.
	env - 애플리케이션에서 사용될 수 있는 환경변수를 나타내는 리소스 타입입니다.
	volume(vol) - 애플리케이션의 영속성있는 저장 공간을 선언하는 리소스 타입입니다.
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
		case resourceType.IsEqual(resource.VOLUME):
			err := getVolume(cmd)
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
		return cmdError.NewCmdError(1, err.Error())
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
		return cmdError.NewCmdError(1, err.Error())
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
		volumeDetail, err := exec.GetVolume(workspaceId, volumeId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		printVolume(*volumeDetail)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("workspace", "w", "", "리소스를 가져올 워크스페이스 아이디")
	getCmd.Flags().StringP("id", "", "", "리소스 아이디")
	getCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n워크스페이스를 가져올때 해당 플래그를 사용하면, 해당 플래그는 무시됩니다.\n리소스 아이디를 사용한다면, 이 플래그는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
}