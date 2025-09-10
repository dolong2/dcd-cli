package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

func CertificateAuthCode(email string, code string) error {
	certificateAuthCodeRequest, err := json.Marshal(request.CertificateAuthCodeRequest{Email: email, Code: code})
	if err != nil {
		return err
	}

	_, err = api.SendPost("/auth/email/certificate", map[string]string{}, map[string]string{}, certificateAuthCodeRequest)
	if err != nil {
		return err
	}

	return nil
}
