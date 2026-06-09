package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

const volumeId contextKey = "volume"

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "볼륨 관련 작업을 식별하기 위한 커맨드",
	Long:  `볼륨에 관련된 작업을 수행하는 커맨드입니다.`,
	Aliases: []string{"vol"},
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 공통적으로 args에 있는 applicationId를 컨텍스트로 전달
		if len(args) > 0 {
			ctx := context.WithValue(cmd.Context(), volumeId, args[0])
			cmd.SetContext(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)

	volumeCmd.PersistentFlags().StringP("workspace", "w", "", "워크스페이스 아이디")
}
