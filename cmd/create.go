package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "command to create a resource",
	Long:  `this command is used to create a resource.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileDirectory, fileErr := cmd.Flags().GetString("file")
		template, templateErr := cmd.Flags().GetString("template")
		if fileErr != nil || templateErr != nil {
			return cmdError.NewCmdError(2, "올바르지 않은 플래그입니다.")
		}
		if fileDirectory == "" && template == "" {
			err := cmdError.NewCmdError(1, "파일 플래그나 템플릿 플래그 중 하나는 입력되어야합니다.")
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
	createCmd.Flags().StringP("template", "t", "", "resource template to json")
	rootCmd.AddCommand(createCmd)
}
