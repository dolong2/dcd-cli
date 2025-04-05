package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type updateMetaData struct {
	ResourceType string `json:"resourceType" yaml:"resourceType"`
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
}

type parsingUpdateMetaData struct {
	Metadata updateMetaData `json:"metadata" yaml:"metadata"`
}

type updateWorkspaceTemplate struct {
	Metadata updateMetaData `json:"metadata" yaml:"metadata"`
}

type updateWorkspaceRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateApplicationTemplate struct {
	Metadata metaData                      `json:"metadata" yaml:"metadata"`
	Spec     updateApplicationSpecTemplate `json:"spec" yaml:"spec"`
}

type updateApplicationSpecTemplate struct {
	GithubUrl       string `json:"githubUrl" yaml:"githubUrl"`
	ApplicationType string `json:"applicationType" yaml:"applicationType"`
	Port            int    `json:"port" yaml:"port"`
	Version         string `json:"version" yaml:"version"`
}

type updateApplicationRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	GithubUrl       string `json:"githubUrl"`
	ApplicationType string `json:"applicationType"`
	Port            int    `json:"port"`
	Version         string `json:"version"`
}

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
	var data parsingUpdateMetaData
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
		var workspace updateWorkspaceTemplate
		err = json.Unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateWorkspaceRequest{Title: workspace.Metadata.Name, Description: workspace.Metadata.Description})
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

		var application updateApplicationTemplate
		err = json.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateApplicationRequest{
			Name:            application.Metadata.Name,
			Description:     application.Metadata.Description,
			GithubUrl:       application.GithubUrl,
			ApplicationType: application.ApplicationType,
			Port:            application.Port,
			Version:         application.Version,
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
	var data parsingUpdateMetaData
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
		var workspace updateWorkspaceTemplate
		err = yaml.Unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateWorkspaceRequest{Title: workspace.Metadata.Name, Description: workspace.Metadata.Description})
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

		var application updateApplicationTemplate
		err = yaml.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := json.Marshal(updateApplicationRequest{
			Name:            application.Metadata.Name,
			Description:     application.Metadata.Description,
			GithubUrl:       application.GithubUrl,
			ApplicationType: application.ApplicationType,
			Port:            application.Port,
			Version:         application.Version,
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
