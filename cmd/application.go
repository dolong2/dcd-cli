package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// context에 값을 담을 때 사용할 전용 커스텀 타입 (충돌 방지)
type contextKey string
const applicationId contextKey = "application"

var applicationCmd = &cobra.Command{
	Use:   "application [applicationId]",
	Short: "애플리케이션 관련 작업을 식별하기 위한 커맨드",
	Long:  `애플리케이션에 관련된 작업을 수행하는 커맨드입니다.`,
	Aliases: []string{"app"},
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 공통적으로 args에 있는 applicationId를 컨텍스트로 전달
		if len(args) > 0 {
			ctx := context.WithValue(cmd.Context(), applicationId, args[0])
			cmd.SetContext(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(applicationCmd)

	applicationCmd.PersistentFlags().StringP("workspace", "w", "", "워크스페이스 아이디")
	applicationCmd.PersistentFlags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n이 플래그를 사용할때 명시한 애플리케이션 아이디는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
}
