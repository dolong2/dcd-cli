package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login in to DCD",
	Long:  `Login to DCD using a email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
	},
}

func init() {
	loginCmd.Flags().BoolP("password", "p", false, "this flag is required to enter password")
	loginCmd.Flags().StringP("email", "e", "", "this flag is used to save user email")
	rootCmd.AddCommand(loginCmd)
}
