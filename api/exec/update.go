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
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

type parsingUpdateMetaData struct {
	Metadata updateMetaData `json:"metadata"`
}

type updateWorkspaceTemplate struct {
	Metadata updateMetaData `json:"metadata"`
}

type updateWorkspaceRequest struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"title"`
	Description  string `json:"description"`
}

type updateApplicationTemplate struct {
	Metadata        metaData `json:"metadata"`
	WorkspaceId     string   `json:"workspaceId"`
	GithubUrl       string   `json:"githubUrl"`
	ApplicationType string   `json:"applicationType"`
	Port            int      `json:"port"`
	Version         string   `json:"version"`
}

type updateApplicationRequest struct {
	Name            string            `json:"title"`
	Description     string            `json:"description"`
	GithubUrl       string            `json:"githubUrl"`
	Env             map[string]string `json:"env"`
	ApplicationType string            `json:"applicationType"`
	Port            int               `json:"port"`
	Version         string            `json:"version"`
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
		return errors.New("invalid file extension")
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

		request, err := json.Marshal(updateWorkspaceRequest{Name: workspace.Metadata.Name, Description: workspace.Metadata.Description})
		if err != nil {
			return err
		}

		_, err = api.SendPut("/workspace/"+resourceId, header, request)
		if err != nil {
			return err
		}
	} else if resourceType == "APPLICATION" {
		var application updateApplicationTemplate
		err := yaml.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := yaml.Marshal(updateApplicationRequest{
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

		_, err = api.SendPatch("/"+application.WorkspaceId+"/application/"+resourceId, header, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New(" this resource type is not supported")
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

		request, err := yaml.Marshal(updateWorkspaceRequest{Name: workspace.Metadata.Name, Description: workspace.Metadata.Description})
		if err != nil {
			return err
		}

		_, err = api.SendPut("/workspace/"+resourceId, header, request)
		if err != nil {
			return err
		}
	} else if resourceType == "APPLICATION" {
		var application updateApplicationTemplate
		err := yaml.Unmarshal(content, &application)
		if err != nil {
			return err
		}

		request, err := yaml.Marshal(updateApplicationRequest{
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

		_, err = api.SendPatch("/"+application.WorkspaceId+"/application/"+resourceId, header, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New(" this resource type is not supported")
	}

	return nil
}
