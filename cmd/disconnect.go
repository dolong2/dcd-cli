package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// disconnectCmd represents the disconnect command
var disconnectCmd = &cobra.Command{
	Use:   "disconnect <domainId>",
	Short: "연결된 도메인을 연결해제하는 커맨드",
	Long: `애플리케이션과 연결된 도메인을 연결해제하는 커맨드입니다.

해당 커맨드가 실행되면 연결되어있던 애플리케이션은 외부로부터 요청을 받을 수 없게됩니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "도메인 아이디가 입력되어야합니다.")
		} else if len(args) > 1 {
			return cmdError.NewCmdError(1, "도매인 아이디만 입력되어야합니다.")
		}

		domainId := args[0]
		err = exec.DisconnectDomain(workspaceId, domainId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)

	disconnectCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
}
