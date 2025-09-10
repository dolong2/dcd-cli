package exec

import (
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

func SignUp(email string, password string, name string) error {
	signUpRequest, err := json.Marshal(request.SignUpRequest{Email: email, Password: password, Name: name})

	if err != nil {
		return err
	}

	_, err = api.SendPost("/auth/signup", map[string]string{}, map[string]string{}, signUpRequest)
	if err != nil {
		return err
	}

	return nil
}
