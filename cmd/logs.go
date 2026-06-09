package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs <applicationId>",
	Short: "애플리케이션의 로그를 조회하는 커맨드",
	Long:  `애플리케이션의 로그를 조회하는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		applicationId, ok := cmd.Context().Value(applicationId).(string)
		if !ok {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
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
	applicationCmd.AddCommand(logsCmd)
}
