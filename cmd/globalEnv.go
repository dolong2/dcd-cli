package cmd

import (
	"github.com/spf13/cobra"
)

// globalEnvCmd represents the globalEnv command
var globalEnvCmd = &cobra.Command{
	Use:   "global-env",
	Short: "전역 환경 변수를 관리하는 커맨드",
	Long:  `워크스페이스의 전역 환경변수를 관리할때 사용하는 커맨드입니다.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(globalEnvCmd)
}
