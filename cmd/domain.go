package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

const domainId contextKey = "domain"

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "도메인 관련 작업을 식별하기 위한 커맨드",
	Long:  `도메인에 관련된 작업을 수행하는 커맨드입니다.`,
	Aliases: []string{"dom"},
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 공통적으로 args에 있는 domainId를 컨텍스트로 전달
		if len(args) > 0 {
			ctx := context.WithValue(cmd.Context(), domainId, args[0])
			cmd.SetContext(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)

	domainCmd.PersistentFlags().StringP("workspace", "w", "", "워크스페이스 아이디")
}
