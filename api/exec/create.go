package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type metaData struct {
	ResourceType string `json:"resourceType" yaml:"resourceType"`
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
}

type parsingMetaData struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

type workspaceTemplate struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

type workspaceRequest struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"title"`
	Description  string `json:"description"`
}

type applicationTemplate struct {
	Metadata        metaData          `json:"metadata" yaml:"metadata"`
	GithubUrl       string            `json:"githubUrl" yaml:"githubUrl"`
	Env             map[string]string `json:"env" yaml:"env"`
	ApplicationType string            `json:"applicationType" yaml:"applicationType"`
	Port            int               `json:"port" yaml:"port"`
	Version         string            `json:"version" yaml:"version"`
	Labels          []string          `json:"labels" yaml:"labels"`
}

type applicationRequest struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	GithubUrl       string            `json:"githubUrl"`
	Env             map[string]string `json:"env"`
	ApplicationType string            `json:"applicationType"`
	Port            int               `json:"port"`
	Version         string            `json:"version"`
	Labels          []string          `json:"labels"`
}

type createWorkspaceResponse struct {
	WorkspaceId string `json:"workspaceId"`
}

type createApplicationResponse struct {
	ApplicationId string `json:"applicationId"`
}

func CreateByPath(fileDirectory string) error {
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	var resourceId = ""
	ext := filepath.Ext(fileDirectory)
	switch ext {
	case ".json":
		resourceId, err = createByJson(content)
		if err != nil {
			return err
		}
	case ".yml", ".yaml":
		resourceId, err = createByYml(content)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid file extension")
	}

	if resourceId != "" {
		err := MapFileToResourceId(fileDirectory, resourceId)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateByTemplate(rawTemplate string) error {
	_, err := createByJson([]byte(rawTemplate))
	if err != nil {
		return err
	}

	return nil
}

func createByJson(content []byte) (string, error) {
	var data parsingMetaData
	err := json.Unmarshal(content, &data)
	if err != nil {
		return "", err
	}

	header := make(map[string]string)
	token, err := GetAccessToken()
	header["Authorization"] = "Bearer " + token
	if err != nil {
		return "", err
	}

	resourceType := data.Metadata.ResourceType
	if resourceType == "WORKSPACE" {
		var workspace workspaceTemplate
		err = json.Unmarshal(content, &workspace)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(workspaceRequest{Name: workspace.Metadata.Name, Description: workspace.Metadata.Description})
		if err != nil {
			return "", err
		}

		response, err := api.SendPost("/workspace", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		createWorkspaceResponse := createWorkspaceResponse{}
		err = json.Unmarshal(response, &createWorkspaceResponse)
		if err != nil {
			return "", err
		}
		return createWorkspaceResponse.WorkspaceId, nil
	} else if resourceType == "APPLICATION" {
		var application applicationTemplate
		err = json.Unmarshal(content, &application)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(applicationRequest{
			Name:            application.Metadata.Name,
			Description:     application.Metadata.Description,
			GithubUrl:       application.GithubUrl,
			Env:             application.Env,
			ApplicationType: application.ApplicationType,
			Port:            application.Port,
			Version:         application.Version,
			Labels:          application.Labels,
		})
		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		response, err := api.SendPost("/"+workspaceId+"/application", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}
		var createApplicationResponse createApplicationResponse
		err = json.Unmarshal(response, &createApplicationResponse)
		if err != nil {
			return "", err
		}
		return createApplicationResponse.ApplicationId, nil
	} else {
		return "", errors.New("지원되지 않는 리소스 타입입니다.")
	}
}

func createByYml(content []byte) (string, error) {
	var data parsingMetaData
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		return "", err
	}

	header := make(map[string]string)
	token, err := GetAccessToken()
	header["Authorization"] = "Bearer " + token
	if err != nil {
		return "", err
	}

	resourceType := data.Metadata.ResourceType
	if resourceType == "WORKSPACE" {
		var workspace workspaceTemplate
		err = yaml.Unmarshal(content, &workspace)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(workspaceRequest{Name: workspace.Metadata.Name, Description: workspace.Metadata.Description})
		if err != nil {
			return "", err
		}

		response, err := api.SendPost("/workspace", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}
		createWorkspaceResponse := createWorkspaceResponse{}
		err = json.Unmarshal(response, &createWorkspaceResponse)
		if err != nil {
			return "", err
		}
		return createWorkspaceResponse.WorkspaceId, nil
	} else if resourceType == "APPLICATION" {
		var application applicationTemplate
		err := yaml.Unmarshal(content, &application)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(applicationRequest{
			Name:            application.Metadata.Name,
			Description:     application.Metadata.Description,
			GithubUrl:       application.GithubUrl,
			Env:             application.Env,
			ApplicationType: application.ApplicationType,
			Port:            application.Port,
			Version:         application.Version,
			Labels:          application.Labels,
		})
		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		response, err := api.SendPost("/"+workspaceId+"/application", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		// 애플리케이션 생성후 애플리케이션 아이디를 반환
		var createApplicationResponse createApplicationResponse
		err = json.Unmarshal(response, &createApplicationResponse)
		if err != nil {
			return "", err
		}
		return createApplicationResponse.ApplicationId, nil
	} else {
		return "", errors.New("지원되지 않는 리소스 타입입니다.")
	}
}
