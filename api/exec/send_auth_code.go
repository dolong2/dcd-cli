package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

func SendAuthCode(email string, usage string) error {
	sendAuthCodeRequest := request.SendAuthCodeRequest{Email: email, Usage: usage}
	rawSendAuthCodeRequest, err := json.Marshal(sendAuthCodeRequest)
	if err != nil {
		return err
	}

	_, err = api.SendPost("/auth/email", map[string]string{}, map[string]string{}, rawSendAuthCodeRequest)
	if err != nil {
		return err
	}

	return nil
}
