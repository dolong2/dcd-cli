package util

import (
	"encoding/json"
	"errors"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ResolveFileExtension(fileDirectory string) (func([]byte, interface{}) (err error), error) {
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