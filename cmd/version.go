package cmd

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version <resourceType>",
	Short: "특정 리소스 타입의 버전을 출력하는 커맨드",
	Long:  `플래그로 입력된 리소스 타입의 버전을 출력합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceType := args[0]
		if resourceType == "" {
			return cmdError.NewCmdError(1, "리소스 타입이 입력되지 않았습니다.")
		}

		versionList, err := exec.GetVersion(resourceType)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		for _, version := range versionList {
			fmt.Println(version)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
