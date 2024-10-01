package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <applicationId> [flags]",
	Short: "a command to execute a command in an application",
	Long: `this command is able to execute a command in an application.
it also can be supported to web socket.
externally this command
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		applicationId := args[0]

		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		command, err := cmd.Flags().GetString("command")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		cmdResult, err := exec.ExecCommand(workspaceId, applicationId, command)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		for _, result := range cmdResult {
			fmt.Println(result)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.Flags().StringP("command", "c", "", "a command to execute in an application")
	execCmd.Flags().StringP("workspace", "w", "", "workspace id")
}
