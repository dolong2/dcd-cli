package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// putGlobalCmd represents the put command
var putGlobalCmd = &cobra.Command{
	Use:   "put",
	Short: "update a exists global env",
	Long:  `this command is used to update a exits global env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId := ""

		if len(args) == 0 {
			var err error = nil
			workspaceId, err = util.GetWorkspaceId(cmd)
			if err != nil {
				return err
			}
		} else {
			workspaceId = args[0]
		}

		application := args[0]
		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both and key and value")
		}
		err := exec.UpdateGlobalEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	globalEnvCmd.AddCommand(putGlobalCmd)

	rmGlobalEnvCmd.Flags().StringP("key", "", "", "select a key to delete")
}
