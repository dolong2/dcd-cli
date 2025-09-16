package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/resource"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <resourceType> <resourceId> [flags]",
	Short: "리소스를 삭제하기 위한 커맨드",
	Long:  `이 커맨드는 리소스를 삭제하기 위해 사용되는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmdError.NewCmdError(1, "리소스 타입과 리소스 아이디가 입력되어야합니다.")
		}

		resourceType := resource.Type(args[0])
		if !resourceType.IsValid() {
			return cmdError.NewCmdError(1, "올바르지 않은 리소스 타입입니다.")
		}

		switch {
		case resourceType.IsEqual(resource.Workspace):
			var workspaceId = args[1]
			err := exec.DeleteWorkspace(workspaceId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		case resourceType.IsEqual(resource.Application):
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			var applicationId = args[1]
			err = exec.DeleteApplication(workspaceId, applicationId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		case resourceType.IsEqual(resource.Env):
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			envId := args[1]

			err = exec.DeleteEnv(workspaceId, envId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		case resourceType.IsEqual(resource.Domain):
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			domainId := args[1]
			err = exec.DeleteDomain(workspaceId, domainId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		case resourceType.IsEqual(resource.VOLUME):
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			volumeId := args[1]
			err = exec.DeleteVolume(workspaceId, volumeId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		default:
			return cmdError.NewCmdError(1, "삭제할 수 없는 리소스 타입입니다.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	deleteCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기 위한 라벨.\nex). -l test-label-1 -l test-label-2")
	deleteCmd.Flags().StringP("application", "a", "", "애플리케이션 아이디")
}
