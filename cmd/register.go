package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "유저 가입 커맨드",
	Long:  `유저 회원가입을 진행하는 커맨드 입니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var email string
		cmd.Print("이메일: ")
		_, err := fmt.Scanf("%s", &email)
		if email == "" || err != nil {
			return cmdError.NewCmdError(1, "올바른 이메일을 입력해주세요")
		}

		err = exec.SendAuthCode(email, "SIGNUP")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

	enterCode:
		cmd.Print("인증 코드: ")
		var authCode string
		_, err = fmt.Scanf("%s", &authCode)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		err = exec.CertificateAuthCode(email, authCode)
		if err != nil {
			cmd.Println(err.Error())
			goto enterCode
		}

		var name string
		cmd.Print("이름: ")
		_, err = fmt.Scanf("%s", &name)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
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
