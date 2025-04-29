package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type metaData struct {
	ResourceType string  `json:"resourceType" yaml:"resourceType"`
	Name         *string `json:"name,omitempty" yaml:"name,omitempty"`
	Description  *string `json:"description,omitempty" yaml:"description,omitempty"`
}

type parsingMetaData struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

type workspaceTemplate struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

type workspaceRequest struct {
	ResourceType string  `json:"resourceType"`
	Name         *string `json:"title"`
	Description  *string `json:"description"`
}

type applicationTemplate struct {
	Metadata metaData                `json:"metadata" yaml:"metadata"`
	Spec     applicationSpecTemplate `json:"spec" yaml:"spec"`
}

type applicationSpecTemplate struct {
	GithubUrl       string   `json:"githubUrl" yaml:"githubUrl"`
	ApplicationType string   `json:"applicationType" yaml:"applicationType"`
	Port            int      `json:"port" yaml:"port"`
	Version         string   `json:"version" yaml:"version"`
	Labels          []string `json:"labels" yaml:"labels"`
}

type applicationRequest struct {
	Name            *string  `json:"name"`
	Description     *string  `json:"description"`
	GithubUrl       string   `json:"githubUrl"`
	ApplicationType string   `json:"applicationType"`
	Port            int      `json:"port"`
	Version         string   `json:"version"`
	Labels          []string `json:"labels"`
}

type envTemplate struct {
	Metadata metaData        `json:"metadata" yaml:"metadata"`
	Spec     envSpecTemplate `json:"spec" yaml:"spec"`
}

type envSpecTemplate struct {
	EnvList       []envListTemplate `json:"envList" yaml:"envList"`
	Labels        []string          `json:"labels" yaml:"labels"`
	ApplicationId *string           `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`
}

type envListTemplate struct {
	Key        string `json:"key" yaml:"key"`
	Value      string `json:"value" yaml:"value"`
	Encryption bool   `json:"encryption" yaml:"encryption"`
}

type envPutRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Encryption bool   `json:"encryption"`
}

type envPutListRequest struct {
	EnvList []envPutRequest `json:"envList"`
}

type globalEnvTemplate struct {
	Metadata metaData              `json:"metadata" yaml:"metadata"`
	Spec     globalEnvSpecTemplate `json:"spec" yaml:"spec"`
}

type globalEnvSpecTemplate struct {
	EnvList     []globalEnvListTemplate `json:"envList" yaml:"envList"`
	WorkspaceId string                  `json:"workspaceId" yaml:"workspaceId"`
}

type globalEnvListTemplate struct {
	Key        string `json:"key" yaml:"key"`
	Value      string `json:"value" yaml:"value"`
	Encryption bool   `json:"encryption" yaml:"encryption"`
}

type globalEnvPutRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Encryption bool   `json:"encryption"`
}

type globalEnvPutListRequest struct {
	EnvList []globalEnvPutRequest `json:"envList"`
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
			GithubUrl:       application.Spec.GithubUrl,
			ApplicationType: application.Spec.ApplicationType,
			Port:            application.Spec.Port,
			Version:         application.Spec.Version,
			Labels:          application.Spec.Labels,
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
	} else if resourceType == "ENV" {
		var envTemplate envTemplate
		err := json.Unmarshal(content, &envTemplate)
		if err != nil {
			return "", err
		}

		var envRequestList []envPutRequest
		for _, template := range envTemplate.Spec.EnvList {
			envRequestList = append(envRequestList, envPutRequest{
				Key:        template.Key,
				Value:      template.Value,
				Encryption: template.Encryption,
			})
		}

		request, err := json.Marshal(envPutListRequest{EnvList: envRequestList})

		if err != nil {
			return "", err
		}

		if envTemplate.Spec.Labels == nil && envTemplate.Spec.ApplicationId == nil {
			return "", errors.New("애플리케이션 아이디 혹은 라벨이 입력되어야함")
		} else if envTemplate.Spec.Labels != nil {
			param := map[string]string{"labels": strings.Join(envTemplate.Spec.Labels, ",")}

			_, err := api.SendPut("/application/env", header, param, request)
			if err != nil {
				return "", err
			}
		} else if envTemplate.Spec.ApplicationId != nil {
			applicationId := *envTemplate.Spec.ApplicationId
			_, err := api.SendPut("/application"+applicationId+"/env", header, map[string]string{}, request)
			if err != nil {
				return "", err
			}
		}

		return "", nil
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
			GithubUrl:       application.Spec.GithubUrl,
			ApplicationType: application.Spec.ApplicationType,
			Port:            application.Spec.Port,
			Version:         application.Spec.Version,
			Labels:          application.Spec.Labels,
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
