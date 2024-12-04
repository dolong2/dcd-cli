package cmd

import (
	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "환경변수를 관리하기 위한 커맨드",
	Long:  `환경변수를 관리하기 위한 커맨드입니다.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
