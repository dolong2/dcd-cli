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

// updateEnvCmd represents the put command
var updateEnvCmd = &cobra.Command{
	Use:   "update",
	Short: "애플리케이션 환경변수를 수정하는 커맨드",
	Long:  `애플리케이션의 환경변수를 수정하는 커맨드 입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "환경변수의 키와 값이 입력되어야합니다.")
		}

		labels, err := cmd.Flags().GetStringArray("label")
		if err != nil {
			return err
		}
		if len(labels) != 0 {
			err := exec.UpdateEnvWithLabel(workspaceId, labels, key, value)
			if err != nil {
				return err
			}
			return nil
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		application := args[0]
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
	Short: "전역 환경변수를 수정하는 커맨드",
	Long:  `존재하는 전역 환경변수를 수정하는 커맨드입니다.`,
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
			return cmdError.NewCmdError(1, "환경변수의 키와 값이 입력되어야합니다.")
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

	updateCmd.Flags().StringP("file", "f", "", "리소스 포맷이 정의된 파일 경로")
	updateCmd.Flags().StringP("template", "", "", "json 포맷으로 정의된 애플리케이션 템플릿")

	envCmd.AddCommand(updateEnvCmd)

	updateEnvCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	updateEnvCmd.Flags().StringP("key", "", "", "수정할 환경변수의 키")
	updateEnvCmd.Flags().StringP("value", "", "", "수정될 환경변수 값")
	updateEnvCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n이 플래그를 사용한다면, 애플리케이션 아이디를 명시할 필요는 없음.\nex). -l test-label-1 -l test-label-2")

	globalEnvCmd.AddCommand(updateGlobalEnvCmd)

	updateGlobalEnvCmd.Flags().StringP("key", "", "", "수정할 환경변수 키")
	updateGlobalEnvCmd.Flags().StringP("value", "", "", "수정될 환경변수 값")
}
