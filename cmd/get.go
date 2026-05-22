package cmd

import (
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/resource"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resourceType> [flags]",
	Short: "리소스를 조회하는 커맨드",
	Long: `이 커맨드는 리소스를 조회하기 위해서 사용됩니다.
조회 가능한 리소스 타입:
	workspace(ws) - 이 리소스 타입은 여러 애플리케이션을 가지고, 작업구역을 나눌때 사용합니다.
	application(app) - 이 리소스 타입은 특정 라이브러리 혹은 프레임워크가 컨테이너에서 동작하게 하는 리소스 타입입니다.
	type(ts) - 애플리케이션의 타입 종류를 나타내는 리소스 타입입니다.
	domain(dom) - 해당 리소스타입은 애플리케이션을 HTTPS로 외부에 공개할때 사용되는 리소스 타입입니다.
	env - 애플리케이션에서 사용될 수 있는 환경변수를 나타내는 리소스 타입입니다.
	volume(vol) - 애플리케이션의 영속성있는 저장 공간을 선언하는 리소스 타입입니다.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "리소스 타입이 입력되어야 합니다.")
		}
		resourceType := resource.Type(args[0])
		if !resourceType.IsValid() {
			return cmdError.NewCmdError(1, "올바르지 않은 리소스 타입입니다.")
		}

		err := util.GetResource(cmd, resourceType)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("workspace", "w", "", "리소스를 가져올 워크스페이스 아이디")
	getCmd.Flags().StringP("id", "", "", "리소스 아이디")
	getCmd.Flags().StringArrayP("label", "l", []string{}, "애플리케이션을 식별하기위한 라벨.\n워크스페이스를 가져올때 해당 플래그를 사용하면, 해당 플래그는 무시됩니다.\n리소스 아이디를 사용한다면, 이 플래그는 무시됩니다.\nex). -l test-label-1 -l test-label-2")
}