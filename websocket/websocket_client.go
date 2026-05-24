package websocket

import (
	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/gorilla/websocket"
	"net/http"
)

var baseUrl string

func Connect(applicationId string) (*websocket.Conn, error) {
	serverUrl := baseUrl + "/application/exec?applicationId=" + applicationId

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
