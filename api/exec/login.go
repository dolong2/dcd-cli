package exec

import (
	"bufio"
	"encoding/json"
	"github.com/dolong2/dcd-cli/api"
	"os"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(tokenRequest *TokenRequest) error {
	rawTokenRequest, err := json.Marshal(tokenRequest)
	if err != nil {
		return err
	}

	rawResult, err := api.SendPost("/auth", map[string]string{}, rawTokenRequest)
	if err != nil {
		return err
	}

	tokenInfoDirectory := "./dcd-info"
	err = os.MkdirAll(tokenInfoDirectory, 0755)
	if err != nil {
		return err
	}

	tokenInfoPath := "./dcd-info/token-info.json"
	file, err := os.Create(tokenInfoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(rawResult))
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
