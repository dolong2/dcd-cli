package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
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

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("file", "f", "", "file path where a resource format is defined")
	updateCmd.Flags().StringP("template", "", "", "resource template to json")
}
