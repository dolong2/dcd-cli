package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "애플리케이션의 로그를 조회하는 커맨드",
	Long:  `애플리케이션의 로그를 조회하는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		applicationId := args[0]
		logs, err := exec.GetLog(workspaceId, applicationId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		for _, log := range logs {
			cmd.Println(log)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	logsCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
}
