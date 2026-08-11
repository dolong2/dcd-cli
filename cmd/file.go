package cmd

import (
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:              "file",
	Short:            "볼륨 내 파일 관련 작업을 식별하기 위한 커맨드",
	Long:             `볼륨 내부의 파일을 관리하는 커맨드입니다.`,
	TraverseChildren: true,
}

func init() {
	volumeCmd.AddCommand(fileCmd)
}
