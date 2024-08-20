package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <resourceId>",
	Short: "command to update a resource",
	Long:  `this command is used to update a resource`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmdError.NewCmdError(1, "must enter both resource type and resource id")
		}

		resourceId := args[0]

		fileDirectory, fileErr := cmd.Flags().GetString("file")
		template, templateErr := cmd.Flags().GetString("template")
		if fileErr != nil || templateErr != nil {
			return cmdError.NewCmdError(2, "invalid flag")
		}
		if fileDirectory == "" && template == "" {
			err := cmdError.NewCmdError(1, "there must be required either file flag or template flag")
			return err
		} else if fileDirectory != "" {
			err := exec.UpdateByPath(resourceId, fileDirectory)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		} else if template != "" {
			err := exec.UpdateByTemplate(resourceId, template)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		}

		return nil
	},
}

// updateEnvCmd represents the put command
var updateEnvCmd = &cobra.Command{
	Use:   "update",
	Short: "update an exists env",
	Long:  `this command is used to update a exits global env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		application := args[0]
		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both and key and value")
		}
		err = exec.UpdateEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

// updateGlobalEnvCmd represents the put command
var updateGlobalEnvCmd = &cobra.Command{
	Use:   "update",
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

		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "this command needs to specify both and key and value")
		}
		err := exec.UpdateGlobalEnv(workspaceId, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("file", "f", "", "file path where a resource format is defined")
	updateCmd.Flags().StringP("template", "", "", "resource template to json")

	envCmd.AddCommand(updateEnvCmd)

	updateEnvCmd.Flags().StringP("key", "", "", "select a key to delete")
	updateEnvCmd.Flags().StringP("workspace", "w", "", "workspace id")

	globalEnvCmd.AddCommand(updateGlobalEnvCmd)

	updateGlobalEnvCmd.Flags().StringP("key", "", "", "select a key to delete")
}
