package websocket

import (
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"os/signal"
)

func Connect(applicationId string) (*websocket.Conn, error) {
	serverUrl := "ws://localhost:8081/application/exec?applicationId=" + applicationId

	header := http.Header{}
	accessToken, err := exec.GetAccessToken()
	if err != nil {
		return nil, err
	}
	header.Add("Authorization", "Bearer "+accessToken)

	conn, _, err := websocket.DefaultDialer.Dial(serverUrl, header)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Close(conn *websocket.Conn) error {
	return conn.Close()
}

func SendMessage(conn *websocket.Conn, message string) error {
	// 인터럽트 신호를 받기 위한 채널
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case <-interrupt:
		fmt.Println("Interrupt signal received. Closing connection...")
		err := conn.Close()
		if err != nil {
			return err
		}
		return nil
	default:
	}
	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func ReadMessage(conn *websocket.Conn) (string, error) {
	_, result, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}

	if string(result) == "cmd start" {
		return "", nil
	}

	return string(result), nil
}
