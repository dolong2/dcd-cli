package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect <domainId> <applicationId> [flags]",
	Short: "도메인을 애플리케이션에 연결하는 커맨드",
	Long: `도메인을 애플리케이션에 연결하기 위해 사용되는 명령입니다.

해당 명령을 수행하면 애플리케이션에 HTTP 트래픽이 도메인을 통해 들어올 수 있게 됩니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		if len(args) < 2 {
			return cmdError.NewCmdError(1, "필요한 변수가 부족합니다.")
		}

		domainId := args[0]
		applicationId := args[1]

		err = exec.ConnectDomain(workspaceId, domainId, applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
}
