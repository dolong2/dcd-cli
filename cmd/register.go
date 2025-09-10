package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register <email> <name>",
	Short: "유저 가입 커맨드",
	Long:  `유저 회원가입을 진행하는 커맨드 입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return cmdError.NewCmdError(1, "이메일과 이름이 입력되지 않음")
		}

		email := args[0]
		name := args[1]
		if email == "" || name == "" {
			return cmdError.NewCmdError(1, "이메일과 이름이 입력되지 않음")
		}

		err := exec.SendAuthCode(email, "SIGNUP")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

	enterCode:
		cmd.Print("인증 코드: ")
		byteCode, err := term.ReadPassword(int(os.Stdin.Fd()))
		cmd.Println()
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		authCode := string(byteCode)
		err = exec.CertificateAuthCode(email, authCode)
		if err != nil {
			cmd.Println(err.Error())
			goto enterCode
		}

		cmd.Print("패스워드: ")
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		cmd.Println()
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		password := string(bytePassword)

		err = exec.SignUp(email, password, name)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
