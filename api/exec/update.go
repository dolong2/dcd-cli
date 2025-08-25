package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/template"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func UpdateByTemplate(resourceId string, rawTemplate string) error {
	err := update(resourceId, []byte(rawTemplate), json.Unmarshal)
	if err != nil {
		return err
	}

	return nil
}

func UpdateByPath(resourceId string, fileDirectory string) error {
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	unmarshal, err := resolveFileExtension(fileDirectory)
	if err != nil {
		return err
	}
	err = update(resourceId, content, unmarshal)

	return nil
}

func UpdateByOnlyPath(fileDirectory string) error {
	resourceId, err := GetResourceIdByFilePath(fileDirectory)

	if err != nil {
		return errors.New("파일이랑 매핑되는 리소스 아이디가 존재하지 않습니다.")
	}

	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	ext := filepath.Ext(fileDirectory)
	switch ext {
	case ".json":
		err = update(resourceId, content, json.Unmarshal)
		if err != nil {
			return err
		}
	case ".yml", ".yaml":
		err = update(resourceId, content, yaml.Unmarshal)
		if err != nil {
			return err
		}
	default:
		return errors.New("지원되지 않는 파일 확장자입니다.")
	}

	return nil
}

func update(resourceId string, content []byte, unmarshal func([]byte, interface{}) (err error)) error {
	var data template.ParsingMetaData
	err := unmarshal(content, &data)
	if err != nil {
		return err
	}

	header := make(map[string]string)
	token, err := GetAccessToken()
	header["Authorization"] = "Bearer " + token
	if err != nil {
		return err
	}

	resourceType := data.Metadata.ResourceType

	switch resourceType {
	case "WORKSPACE":
		var workspace template.WorkspaceTemplate
		err = unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		updateWorkspaceRequest, err := workspace.ToRequest()
		if err != nil {
			return err
		}
		request, err := json.Marshal(updateWorkspaceRequest)
		if err != nil {
			return err
		}

		_, err = api.SendPut("/workspace/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	case "APPLICATION":
		workspaceId, err := getWorkspaceId()
		if err != nil {
			return err
		}

		var application template.ApplicationTemplate
		err = unmarshal(content, &application)
		if err != nil {
			return err
		}

		updateApplicationRequest, err := application.ToRequest()
		if err != nil {
			return err
		}
		request, err := json.Marshal(updateApplicationRequest)
		if err != nil {
			return err
		}

		_, err = api.SendPut("/"+workspaceId+"/application/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	case "ENV":
		workspaceId, err := getWorkspaceId()
		if err != nil {
			return err
		}

		var env template.EnvTemplate
		err = unmarshal(content, &env)
		if err != nil {
			return err
		}

		updateEnvRequest := env.ToRequest()

		request, err := json.Marshal(updateEnvRequest)
		if err != nil {
			return err
		}

		_, err = api.SendPut("/"+workspaceId+"/env/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	default:
		return errors.New("지원되지 않는 리소스 타입입니다")
	}

	return nil
}
