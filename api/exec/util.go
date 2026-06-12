package exec

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
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
	accessTokenExp, err := time.ParseInLocation(timeFormat, raw["accessTokenExp"].(string), time.Local)
	if err != nil {
		return "", errors.New("RefreshTokenExp 파싱 중 오류 발생")
	}
	refreshTokenExp, err := time.ParseInLocation(timeFormat, raw["refreshTokenExp"].(string), time.Local)
	if err != nil {
		return "", errors.New("RefreshTokenExp 파싱 중 오류 발생")
	}

	// AccessToken이 만료되었을때 토큰 재발급
	now := time.Now().Local()
	if now.After(accessTokenExp) {
		err := ReissueToken(raw["refreshToken"].(string))
		if err != nil {
			return "", err
		}
		goto getTokenInfo
	} else if now.After(refreshTokenExp) {
		return "", errors.New("로그인 정보가 만료되었습니다.\n다시 로그인해주세요.")
	}

	tokenInfo := tokenResponse{
		AccessToken:     raw["accessToken"].(string),
		RefreshToken:    raw["refreshToken"].(string),
		AccessTokenExp:  accessTokenExp,
		RefreshTokenExp: refreshTokenExp,
	}
	return tokenInfo.AccessToken, nil
}

func resolveFileExtension(fileDirectory string) (func([]byte, interface{}) (err error), error) {
	ext := filepath.Ext(fileDirectory)

	switch ext {
	case ".json":
		return json.Unmarshal, nil
	case ".yml", ".yaml":
		return yaml.Unmarshal, nil
	default:
		return nil, errors.New("지원되지 않는 파일 확장자입니다.")
	}
}
