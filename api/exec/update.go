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

func UpdateByTemplate(workspaceId string, rawTemplate string) error {
	err := updateByJson(workspaceId, []byte(rawTemplate))
	if err != nil {
		return err
	}

	return nil
}

func UpdateByPath(workspaceId string, fileDirectory string) error {
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	ext := filepath.Ext(fileDirectory)
	switch ext {
	case ".json":
		err = updateByJson(workspaceId, content)
		if err != nil {
			return err
		}
	case ".yml", ".yaml":
		err = updateByYml(workspaceId, content)
		if err != nil {
			return err
		}
	default:
		return errors.New("지원되지 않는 파일 확장자입니다.")
	}

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
		err = updateByJson(resourceId, content)
		if err != nil {
			return err
		}
	case ".yml", ".yaml":
		err = updateByYml(resourceId, content)
		if err != nil {
			return err
		}
	default:
		return errors.New("지원되지 않는 파일 확장자입니다.")
	}

	return nil
}

func updateByJson(resourceId string, content []byte) error {
	var data template.ParsingMetaData
	err := json.Unmarshal(content, &data)
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
	if resourceType == "WORKSPACE" {
		var workspace template.WorkspaceTemplate
		err = json.Unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		err = workspace.ValidateMetadata()
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateWorkspaceRequest{Title: *workspace.Metadata.Name, Description: *workspace.Metadata.Description})
		if err != nil {
			return err
		}

		_, err = api.SendPut("/workspace/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else if resourceType == "APPLICATION" {
		workspaceId, err := getWorkspaceId()
		if err != nil {
			return err
		}

		var application template.ApplicationTemplate
		err = json.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateApplicationRequest{
			Name:            *application.Metadata.Name,
			Description:     *application.Metadata.Description,
			GithubUrl:       application.Spec.GithubUrl,
			ApplicationType: application.Spec.ApplicationType,
			Port:            application.Spec.Port,
			Version:         application.Spec.Version,
		})
		if err != nil {
			return err
		}

		_, err = api.SendPut("/"+workspaceId+"/application/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New("지원되지 않는 리소스 타입입니다.")
	}

	return nil
}

func updateByYml(resourceId string, content []byte) error {
	var data template.ParsingMetaData
	err := yaml.Unmarshal(content, &data)
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
	if resourceType == "WORKSPACE" {
		var workspace template.WorkspaceTemplate
		err = yaml.Unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		err = workspace.ValidateMetadata()
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateWorkspaceRequest{Title: *workspace.Metadata.Name, Description: *workspace.Metadata.Description})
		if err != nil {
			return err
		}

		_, err = api.SendPut("/workspace/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else if resourceType == "APPLICATION" {
		workspaceId, err := getWorkspaceId()
		if err != nil {
			return err
		}

		var application template.ApplicationTemplate
		err = yaml.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateApplicationRequest{
			Name:            *application.Metadata.Name,
			Description:     *application.Metadata.Description,
			GithubUrl:       application.Spec.GithubUrl,
			ApplicationType: application.Spec.ApplicationType,
			Port:            application.Spec.Port,
			Version:         application.Spec.Version,
		})
		if err != nil {
			return err
		}

		_, err = api.SendPut("/"+workspaceId+"/application/"+resourceId, header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New("지원되지 않는 리소스 타입입니다")
	}

	return nil
}
