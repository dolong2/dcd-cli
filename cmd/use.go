package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use <workspaceId>",
	Short: "사용하는 워크스페이스를 지정하기 위한 커맨드",
	Long:  `작업하기 위해서 사용할 워크스페이스를 지정하기 위한 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmdError.NewCmdError(1, "워크스페이스 아이디가 입력되어야합니다.")
		}
		workspaceId := args[0]
		if workspaceId == "" || workspaceId == " " {
			return cmdError.NewCmdError(1, "워크스페이스 아이디가 입력되어야합니다.")
		}

		workspace, err := exec.GetWorkspace(workspaceId)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		err = util.SaveWorkspaceInfo(*workspace)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
