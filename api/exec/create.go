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

func CreateByPath(fileDirectory string) error {
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	ext := filepath.Ext(fileDirectory)
	switch ext {
	case ".json":
		err = createByJson(content)
		if err != nil {
			return err
		}
	case ".yml", ".yaml":
		err = createByYml(content)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid file extension")
	}

	return nil
}

func CreateByTemplate(rawTemplate string) error {
	err := createByJson([]byte(rawTemplate))
	if err != nil {
		return err
	}

	return nil
}

func createByJson(content []byte) error {
	var data parsingMetaData
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
		var workspace workspaceTemplate
		err = json.Unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		request, err := json.Marshal(workspaceRequest{Name: workspace.Metadata.Name, Description: workspace.Metadata.Description})
		if err != nil {
			return err
		}

		_, err = api.SendPost("/workspace", header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else if resourceType == "APPLICATION" {
		var application applicationTemplate
		err := json.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := json.Marshal(applicationRequest{
			Name:            application.Metadata.Name,
			Description:     application.Metadata.Description,
			GithubUrl:       application.GithubUrl,
			Env:             application.Env,
			ApplicationType: application.ApplicationType,
			Port:            application.Port,
			Version:         application.Version,
		})
		if err != nil {
			return err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return err
		}

		_, err = api.SendPost("/"+workspaceId+"/application", header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New(" this resource type is not supported")
	}

	return nil
}

func createByYml(content []byte) error {
	var data parsingMetaData
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
		var workspace workspaceTemplate
		err = yaml.Unmarshal(content, &workspace)
		if err != nil {
			return err
		}

		request, err := json.Marshal(workspaceRequest{Name: workspace.Metadata.Name, Description: workspace.Metadata.Description})
		if err != nil {
			return err
		}

		_, err = api.SendPost("/workspace", header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else if resourceType == "APPLICATION" {
		var application applicationTemplate
		err := yaml.Unmarshal(content, &application)
		if err != nil {
			return err
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
			return err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return err
		}

		_, err = api.SendPost("/"+workspaceId+"/application", header, map[string]string{}, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New(" this resource type is not supported")
	}

	return nil
}
