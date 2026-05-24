package websocket

import (
	"net/http"
	"time"

	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/gorilla/websocket"
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
	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		return err
	}
    err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
    if err != nil {
        return err
    }
    err = conn.Close()
    return err
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
