package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <applicationId> [flags]",
	Short: "use to add an application env",
	Long:  `this command can be used to add a env to an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId, err := util.GetWorkspaceId(cmd)
		if err != nil {
			return err
		}

		key, keyErr := cmd.Flags().GetString("key")
		value, valueErr := cmd.Flags().GetString("value")
		if key == "" || value == "" || keyErr != nil || valueErr != nil {
			return cmdError.NewCmdError(1, "환경변수의 키와 값이 전부 입력되어야합니다.")
		}

		labels, err := cmd.Flags().GetStringArray("label")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		// label을 하나라도 받았을때 애플리케이션 id를 사용하지 않고, 라벨로 환경변수를 추가
		if len(labels) != 0 {
			err := exec.AddEnvWithLabels(workspaceId, labels, key, value)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			return nil
		}

		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		application := args[0]
		err = exec.AddEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

// addGlobalEnvCmd represents the add command
var addGlobalEnvCmd = &cobra.Command{
	Use:   "add <workspaceId> [flags]",
	Short: "use to add a global env",
	Long:  `this command can be used to add a global env to an application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceId := ""

		if len(args) == 0 {
			var err error = nil
			workspaceId, err = util.GetWorkspaceId(cmd)
			if err != nil {
				return err
			}
		} else {
			workspaceId = args[0]
		}

		application := args[0]
		key, existsKey := cmd.Flags().GetString("key")
		value, existsValue := cmd.Flags().GetString("value")
		if key == "" || value == "" || existsKey != nil || existsValue != nil {
			return cmdError.NewCmdError(1, "환경변수의 키와 값이 전부 입력되어야합니다.")
		}
		err := exec.AddEnv(workspaceId, application, key, value)
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}
		return nil
	},
}

func init() {
	envCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("key", "k", "", "environment key")
	addCmd.Flags().StringP("value", "v", "", "environment value")
	addCmd.Flags().StringP("workspace", "w", "", "workspace id")
	addCmd.Flags().StringArrayP("label", "l", []string{}, "select labels for applications.\nif use this flag, you are no need to use application id.\nex). -l test-label-1 -l test-label-2")
	globalEnvCmd.AddCommand(addGlobalEnvCmd)
	addGlobalEnvCmd.Flags().StringP("key", "k", "", "environment key")
	addGlobalEnvCmd.Flags().StringP("value", "v", "", "environment value")
}
