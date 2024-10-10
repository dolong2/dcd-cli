package exec

import (
	"encoding/json"
	"errors"
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

func getWorkspaceId() (string, error) {
	rawWorkspaceInfo, err := os.ReadFile("./dcd-info/workspace-info.json")
	if err != nil {
		return "", errors.New("not found workspace info")
	}

	var workspaceInfo map[string]interface{}

	err = json.Unmarshal(rawWorkspaceInfo, &workspaceInfo)
	if err != nil {
		return "", errors.New("invalid workspace info")
	}

	workspaceId := workspaceInfo["workspaceId"].(string)

	return workspaceId, nil
}
