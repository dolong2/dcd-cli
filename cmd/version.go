package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version <resourceType>",
	Short: "특정 리소스 타입의 버전을 출력하는 커맨드",
	Long:  `플래그로 입력된 리소스 타입의 버전을 출력합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return cmdError.NewCmdError(1, "리소스 타입이 입력되지 않았습니다.")
		}
		resourceType := args[0]

		versionList, err := exec.GetVersion(resourceType)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{resourceType + " Version List"})

		for _, versionValue := range versionList {
			table.Append([]string{versionValue})
		}

		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT})
		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
