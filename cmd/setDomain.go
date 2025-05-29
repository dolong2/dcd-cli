package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// setDomainCmd represents the setDomain command
var setDomainCmd = &cobra.Command{
	Use:   "set-domain <applicationId> [flags]",
	Short: "애플리케이션을 외부 도메인이랑 연결하는 커맨드",
	Long:  `특정 애플리케이션에 도메인 이름을 설정해서 외부에 공개하는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		applicationId := args[0]

		domainName, err := cmd.Flags().GetString("name")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		if domainName == "" {
			return cmdError.NewCmdError(1, "도메인 이름이 입력되어야합니다.")
		}

		err = exec.SetDomain(workspaceId, applicationId, domainName)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setDomainCmd)

	setDomainCmd.Flags().StringP("name", "n", "", "서브 도메인 이름")
	setDomainCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
}
