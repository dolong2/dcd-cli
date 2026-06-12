package util

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

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

	fileKey, err := getFileKey(fileDirectory)
	if err != nil {
		return err
	}

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
	data[fileKey] = resourceId

	// 수정된 맵을 JSON으로 마셜링
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// JSON 파일에 저장
	tmpPath := resourceMappingInfoPath + ".tmp"
	if err := os.WriteFile(tmpPath, updatedJSON, 0644); err != nil {
	   return err
	}
	if err := os.Rename(tmpPath, resourceMappingInfoPath); err != nil {
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

	fileKey, err := getFileKey(fileDirectory)
	if err != nil {
		return "", err
	}
	resourceId := data[fileKey]

	if resourceId == "" {
		return "", errors.New("해당 템플릿에 매핑된 리소스 아이디를 찾을 수 없음")
	}

	return resourceId, nil
}