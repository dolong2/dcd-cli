package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "리소스 생성을 위한 커맨드",
	Long:  `이 커맨드는 리소스 생성을 위해 생성될 수 있습니다.`,
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
		fmt.Println("리소스 생성 완료!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("file", "f", "", "리소스 포맷이 정의된 파일 경로")
	createCmd.Flags().StringP("template", "t", "", "json 포맷인 리소스 템플릿")
}
