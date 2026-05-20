package cmd

import (
	"bufio"
	"github.com/dolong2/dcd-cli/api/exec"
	cmdError "github.com/dolong2/dcd-cli/cmd/err"
	"github.com/dolong2/dcd-cli/cmd/util"
	"github.com/dolong2/dcd-cli/websocket"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <applicationId> [flags]",
	Short: "애플리케이션 내부에서 커맨드를 실행하기 위한 커멘드",
	Long: `이 커맨드는 애플리케이션에 커맨드를 실행할 수 있는 커맨드입니다.
웹소켓을 통해서 애플리케이션 내부에 접근할 수 있습니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmdError.NewCmdError(1, "애플리케이션 아이디가 입력되어야합니다.")
		}
		applicationId := args[0]

		ws, err := cmd.Flags().GetBool("ws")
		if err != nil {
			return cmdError.NewCmdError(1, err.Error())
		}

		if ws {
			conn, err := websocket.Connect(applicationId)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
			}
			defer websocket.Close(conn)

			// 인터럽트 신호를 받기 위한 채널
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)

			// 에러를 전송받기 위한 채널
			errChan := make(chan error, 1)

			// [1] 메시지 수신을 위한 독립적인 고루틴 실행
			go func() {
				for {
					message, err := websocket.ReadMessage(conn)
					if err != nil {
						errChan <- err
						return
					}
					// 서버가 빈 메시지를 주든 아니든, 오는 대로 바로 출력
					cmd.Print(message) 
				}
			}()

			// [2] 메시지 송신 루프 (메인 흐름)
			go func() {
				reader := bufio.NewReader(os.Stdin)
				for {
					input, _, err := reader.ReadLine()
					if err != nil {
						errChan <- err
						return
					}

					err = websocket.SendMessage(conn, string(input))
					if err != nil {
						errChan <- err
						return
					}
				}
			}()

			// [3] 인터럽트 및 에러 대기 제어
			select {
			case <-interrupt:
				return nil
			case err := <-errChan:
				return cmdError.NewCmdError(1, err.Error())
			}
		} else {
			workspaceId, err := util.GetWorkspaceId(cmd)
			if err != nil {
				return cmdError.NewCmdError(1, err.Error())
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
				cmd.Println(result)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.Flags().StringP("command", "c", "", "애플리케이션에 실행할 명령")
	execCmd.Flags().StringP("workspace", "w", "", "워크스페이스 아이디")
	execCmd.Flags().Bool("ws", false, "웹소켓 사용 여부")
}
