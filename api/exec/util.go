package exec

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
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
		return "", errors.New("워크스페이스 정보를 찾을수없습니다.")
	}

	var workspaceInfo map[string]interface{}

	err = json.Unmarshal(rawWorkspaceInfo, &workspaceInfo)
	if err != nil {
		return "", errors.New("워크스페이스 정보가 옳바르지 않습니다.")
	}

	workspaceId := workspaceInfo["workspaceId"].(string)

	return workspaceId, nil
}

func MapFileToResourceId(fileDirectory string, resourceId string) error {
	if resourceId == "" || fileDirectory == "" {
		return errors.New("리소스 아이디와 파일경로는 필수적으로 입력되어야합니다.")
	}

	resourceMappingInfoPath := "./dcd-info/resource-mapping-info.json"

	// 디렉토리가 없으면 생성
	resourceMappingDir := filepath.Dir(resourceMappingInfoPath)
	if err := os.MkdirAll(resourceMappingDir, 0755); err != nil {
		return errors.New("dcd-info 정보 디렉토리를 생성하는데 실패했습니다.")
	}

	fileName := filepath.Base(fileDirectory)
	// JSON 파일 읽기 또는 파일이 없을 경우 새로 생성
	file, err := os.ReadFile(resourceMappingInfoPath)
	var data map[string]string

	if os.IsNotExist(err) {
		// 파일이 없을 경우 초기 맵을 생성
		data = make(map[string]string)
	} else if err != nil {
		return err
	} else {
		// 파일이 존재하는 경우 JSON 데이터를 언마셜링
		if err := json.Unmarshal(file, &data); err != nil {
			return err
		}
	}

	// 새로운 key:value 쌍 추가
	data[fileName] = resourceId

	// 수정된 맵을 JSON으로 마셜링
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// JSON 파일에 저장
	if err := os.WriteFile(resourceMappingInfoPath, updatedJSON, 0644); err != nil {
		return err
	}

	return nil
}

func GetResourceIdByFilePath(fileDirectory string) (string, error) {
	// JSON 파일 경로
	filePath := "./dcd-info/resource-mapping-info.json"

	// JSON 파일 읽기
	resourceMappingInfo, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// JSON 데이터 언마샬링
	var data map[string]string
	if err := json.Unmarshal(resourceMappingInfo, &data); err != nil {
		return "", err
	}

	templateName := filepath.Base(fileDirectory)
	resourceId := data[templateName]

	if resourceId == "" {
		return "", errors.New("해당 템플릿에 매핑된 리소스 아이디를 찾을 수 없음")
	}

	return resourceId, nil
}
