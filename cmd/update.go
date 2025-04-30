package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <resourceId>",
	Short: "리소스를 수정하는 커맨드",
	Long:  `리소스 정보를 수정하는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileDirectory, fileErr := cmd.Flags().GetString("file")
		template, templateErr := cmd.Flags().GetString("template")
		if fileErr != nil || templateErr != nil {
			return cmdError.NewCmdError(2, "올바르지 않은 플래그입니다.")
		}

		// args가 없다면 파일명에 매핑된 리소스 아이디를 가져오는 메서드 호출
		if len(args) < 1 {
			if fileDirectory == "" {
				return cmdError.NewCmdError(1, "리소스 아이디는 파일 플래그를 사용할때만 생략할 수 있습니다.")
			}
			err := exec.UpdateByOnlyPath(fileDirectory)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			return nil
		}

		resourceId := args[0]

		if fileDirectory == "" && template == "" {
			err := cmdError.NewCmdError(1, "파일 플래그나 템플릿 플래그가 있어야합니다.")
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

	updateCmd.Flags().StringP("file", "f", "", "리소스 포맷이 정의된 파일 경로")
	updateCmd.Flags().StringP("template", "", "", "json 포맷으로 정의된 애플리케이션 템플릿")
}
