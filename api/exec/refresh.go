package exec

import (
	"bufio"
	"github.com/dolong2/dcd-cli/api"
	"os"
)

func ReissueToken(refreshToken string) error {
	header := map[string]string{}
	header["RefreshToken"] = refreshToken
	rawResult, err := api.SendPost("/auth", header, []byte(""))
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
