package cmd

import (
	"bufio"
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/dolong2/dcd-cli/websocket"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <applicationId> [flags]",
	Short: "a command to execute a command in an application",
	Long: `this command is able to execute a command in an application.
it also can be supported to web socket.
externally this command
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "must specify applicationId")
		}
		applicationId := args[0]

		ws, err := cmd.Flags().GetBool("ws")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		if ws {
			workingDir := "/"
			conn, err := websocket.Connect(applicationId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			// 인터럽트 신호를 받기 위한 채널
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)

			// 사용자 입력을 기다리며 메시지를 서버로 전송
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print(workingDir + " > ")
				input, _, _ := reader.ReadLine() // 사용자 입력 받기

				// 인터럽트 신호 처리 (Ctrl+C)
				select {
				case <-interrupt:
					err := websocket.Close(conn)
					if err != nil {
						return cmdError.NewCmdError(1, err.Error())
					}
					return nil
				default:
				}

				err := websocket.SendMessage(conn, string(input))
				if err != nil {
					return cmdError.NewCmdError(1, err.Error())
				}

				for {
					workingDirPrefix := "current dir = "
					endPrefix := "cmd end"
					message, err := websocket.ReadMessage(conn)
					if err != nil {
						return cmdError.NewCmdError(1, err.Error())
					}
					if strings.HasPrefix(message, workingDirPrefix) {
						workingDir = strings.TrimPrefix(message, workingDirPrefix)
						continue
					} else if strings.HasPrefix(message, endPrefix) {
						break
					}

					fmt.Print(message)
				}

				fmt.Println("")
			}
		} else {
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return err
			}

			command, err := cmd.Flags().GetString("command")
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			cmdResult, err := exec.ExecCommand(workspaceId, applicationId, command)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}

			for _, result := range cmdResult {
				fmt.Println(result)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.Flags().StringP("command", "c", "", "a command to execute in an application")
	execCmd.Flags().StringP("workspace", "w", "", "workspace id")
	execCmd.Flags().Bool("ws", false, "apply websocket")
}
