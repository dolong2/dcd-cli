package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "DCD에서 로그아웃하는 커맨드",
	Long:  `로그인되어있는 정보를 로그아웃시키는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := exec.Logout()
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
