package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use <workspaceId>",
	Short: "command to specify which workspace to primarily use",
	Long:  `this command can be used to specify which workspace to primarily use.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmdError.NewCmdError(1, "must be specify workspaceId")
		}
		workspaceId := args[0]
		if workspaceId == "" || workspaceId == " " {
			return cmdError.NewCmdError(1, "must be specify workspaceId")
		}

		workspace, err := exec.GetWorkspace(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		err = util.SaveWorkspaceInfo(*workspace)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
