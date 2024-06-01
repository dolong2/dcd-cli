package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "command to create a resource",
	Long:  `this command is used to create a resource.`,
	Example: `dcd-cli create [flags]
	resource type is optional
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileDirectory, fileErr := cmd.Flags().GetString("file")
		template, templateErr := cmd.Flags().GetString("template")
		if fileErr != nil || templateErr != nil {
			return cmdError.NewCmdError(2, "invalid flag")
		}
		if fileDirectory == "" && template == "" {
			err := cmdError.NewCmdError(1, "there must be required either file flag or template flag")
			return err
		} else if fileDirectory != "" {
			err := exec.CreateByPath(fileDirectory)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		} else if template != "" {
			err := exec.CreateByTemplate(template)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
		}
		fmt.Println("Successfully created resource!")
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("file", "f", "", "file path where a resource format is defined")
	createCmd.Flags().StringP("template", "", "", "resource template to json")
	rootCmd.AddCommand(createCmd)
}
