package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login in to DCD",
	Long:  `Login to DCD using a email and password`,
	RunE: func(cmd *cobra.Command, args []string) error {
		email, emailErr := cmd.Flags().GetString("email")
		existsPassword, passwordErr := cmd.Flags().GetBool("password")
		if emailErr != nil || passwordErr != nil {
			return cmdError.NewCmdError(1, "invalid flag")
		}
		password := ""
		if existsPassword {
			fmt.Print("Enter password: ")
			bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			password = string(bytePassword)
		}

		tokenRequest := exec.TokenRequest{Email: email, Password: password}
		err := exec.Login(&tokenRequest)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	loginCmd.Flags().BoolP("password", "p", false, "this flag is required to enter password")
	loginCmd.Flags().StringP("email", "e", "", "this flag is used to save user email")
	rootCmd.AddCommand(loginCmd)
}
