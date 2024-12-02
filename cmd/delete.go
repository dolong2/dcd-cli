package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <resourceType> <resourceId> [flags]",
	Short: "command to delete an resource",
	Long:  `this command will delete an resource.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmdError.NewCmdError(1, "리소스 타입과 리소스 아이디가 입력되어야합니다.")
		}

		resourceType := args[0]
		if resourceType == "workspace" {
			var workspaceId = args[1]
			err := exec.DeleteWorkspace(workspaceId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		} else if resourceType == "application" {
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return err
			}

			var applicationId = args[1]
			err = exec.DeleteApplication(workspaceId, applicationId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		} else {
			return cmdError.NewCmdError(1, "옳바르지 않은 리소스 타입입니다.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("workspace", "w", "", "workspace id")
}
