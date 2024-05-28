package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dolong2/dcd-cli/api/exec"
	"os"
	"time"
)

func GetAccessToken() (string, error) {
	type tokenResponse struct {
		AccessToken     string    `json:"accessToken"`
		RefreshToken    string    `json:"refreshToken"`
		AccessTokenExp  time.Time `json:"accessTokenExp"`
		RefreshTokenExp time.Time `json:"refreshTokenExp"`
	}

	const timeFormat = "2006-01-02T15:04:05"
getTokenInfo:
	tokenInfoFile, err := os.Open("./dcd-info/token-info.json")
	if err != nil {
		return "", err
	}
	defer tokenInfoFile.Close()

	decoder := json.NewDecoder(tokenInfoFile)
	var raw map[string]interface{}
	err = decoder.Decode(&raw)
	if err != nil {
		return "", err
	}

	// 커스텀 시간 형식 파싱
	accessTokenExp, err := time.Parse(timeFormat, raw["accessTokenExp"].(string))
	if err != nil {
		fmt.Println("AccessTokenExp 파싱 중 오류 발생:", err)
		return "", err
	}
	refreshTokenExp, err := time.Parse(timeFormat, raw["refreshTokenExp"].(string))
	if err != nil {
		fmt.Println("RefreshTokenExp 파싱 중 오류 발생:", err)
		return "", err
	}

	// AccessToken이 만료되었을때 토큰 재발급
	now := time.Now()
	if now.After(accessTokenExp) {
		err := exec.ReissueToken(raw["refreshToken"].(string))
		if err != nil {
			return "", err
		}
		goto getTokenInfo
	} else if now.After(refreshTokenExp) {
		return "", errors.New("login info is expired.\nplease retry login.")
	}

	tokenInfo := tokenResponse{
		AccessToken:     raw["accessToken"].(string),
		RefreshToken:    raw["refreshToken"].(string),
		AccessTokenExp:  accessTokenExp,
		RefreshTokenExp: refreshTokenExp,
	}
	return tokenInfo.AccessToken, nil
}
