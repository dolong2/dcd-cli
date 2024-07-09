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
			return cmdError.NewCmdError(1, "must enter resource type and resource id")
		}
		if args[0] == "" {
			return cmdError.NewCmdError(1, "must specify a resource type")
		}

		resourceType := args[0]
		if resourceType == "workspace" {
			if args[1] == "" {
				return cmdError.NewCmdError(1, "must specify workspace id")
			}

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

			if args[1] == "" {
				return cmdError.NewCmdError(1, "must specify application id")
			}

			var applicationId = args[1]
			err = exec.DeleteApplication(workspaceId, applicationId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("workspace", "w", "", "workspace id")
}
