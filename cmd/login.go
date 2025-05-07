package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login [flags]",
	Short: "DCD에 로그인하는 커맨드",
	Long:  `이메일과 패스워드를 사용해서 DCD에 로그인하는 커맨드입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		email, emailErr := cmd.Flags().GetString("email")
		existsPassword, passwordErr := cmd.Flags().GetBool("password")
		if email == "" || emailErr != nil || passwordErr != nil {
			return cmdError.NewCmdError(1, "올바르지 않은 플래그입니다.")
		}
		password := ""
		if existsPassword {
			cmd.Print("Enter password: ")
			bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
			cmd.Println()
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			password = string(bytePassword)
		}

		err := exec.Login(email, password)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().BoolP("password", "p", false, "패스워드 사용여부")
	loginCmd.Flags().StringP("email", "e", "", "유저 이메일")
}
