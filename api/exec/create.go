package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
	"github.com/dolong2/dcd-cli/api/exec/template"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

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
	var data template.ParsingMetaData
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
		var workspace template.WorkspaceTemplate
		err = json.Unmarshal(content, &workspace)
		if err != nil {
			return "", err
		}

		err = workspace.ValidateMetadata()
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(workspace.ToRequest())
		if err != nil {
			return "", err
		}

		result, err := api.SendPost("/workspace", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		createWorkspaceResponse := response.CreateWorkspaceResponse{}
		err = json.Unmarshal(result, &createWorkspaceResponse)
		if err != nil {
			return "", err
		}
		return createWorkspaceResponse.WorkspaceId, nil
	} else if resourceType == "APPLICATION" {
		var application template.ApplicationTemplate
		err = json.Unmarshal(content, &application)
		if err != nil {
			return "", err
		}

		err = application.ValidateMetadata()
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(application.ToRequest())
		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		result, err := api.SendPost("/"+workspaceId+"/application", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}
		createApplicationResponse := response.CreateApplicationResponse{}
		err = json.Unmarshal(result, &createApplicationResponse)
		if err != nil {
			return "", err
		}
		return createApplicationResponse.ApplicationId, nil
	} else if resourceType == "ENV" {
		var envTemplate template.EnvTemplate
		err := json.Unmarshal(content, &envTemplate)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(envTemplate.ToRequest())

		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		if envTemplate.Spec.Labels == nil && envTemplate.Spec.ApplicationId == nil {
			return "", errors.New("애플리케이션 아이디 혹은 라벨이 입력되어야함")
		} else if envTemplate.Spec.Labels != nil {
			param := map[string]string{"labels": strings.Join(envTemplate.Spec.Labels, ",")}

			_, err := api.SendPut("/"+workspaceId+"/application/env", header, param, request)
			if err != nil {
				return "", err
			}
		} else if envTemplate.Spec.ApplicationId != nil {
			applicationId := *envTemplate.Spec.ApplicationId
			_, err := api.SendPut("/"+workspaceId+"/application/"+applicationId+"/env", header, map[string]string{}, request)
			if err != nil {
				return "", err
			}
		}

		return "", nil
	} else if resourceType == "GLOBAL_ENV" || resourceType == "GE" {
		var globalEnvTemplate template.GlobalEnvTemplate
		err := json.Unmarshal(content, &globalEnvTemplate)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(globalEnvTemplate.ToRequest())

		if err != nil {
			return "", err
		}

		_, err = api.SendPut("/workspace/"+globalEnvTemplate.Spec.WorkspaceId+"/env", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		return "", nil
	} else {
		return "", errors.New("지원되지 않는 리소스 타입입니다.")
	}
}

func createByYml(content []byte) (string, error) {
	var data template.ParsingMetaData
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
		var workspace template.WorkspaceTemplate
		err = yaml.Unmarshal(content, &workspace)
		if err != nil {
			return "", err
		}

		err = workspace.ValidateMetadata()
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(workspace.ToRequest())
		if err != nil {
			return "", err
		}

		result, err := api.SendPost("/workspace", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}
		createWorkspaceResponse := response.CreateWorkspaceResponse{}
		err = json.Unmarshal(result, &createWorkspaceResponse)
		if err != nil {
			return "", err
		}
		return createWorkspaceResponse.WorkspaceId, nil
	} else if resourceType == "APPLICATION" {
		var application template.ApplicationTemplate
		err := yaml.Unmarshal(content, &application)
		if err != nil {
			return "", err
		}

		err = application.ValidateMetadata()
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(application.ToRequest())
		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		result, err := api.SendPost("/"+workspaceId+"/application", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		// 애플리케이션 생성후 애플리케이션 아이디를 반환
		createApplicationResponse := response.CreateApplicationResponse{}
		err = json.Unmarshal(result, &createApplicationResponse)
		if err != nil {
			return "", err
		}
		return createApplicationResponse.ApplicationId, nil
	} else if resourceType == "ENV" {
		var envTemplate template.EnvTemplate
		err := yaml.Unmarshal(content, &envTemplate)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(envTemplate.ToRequest())

		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		if envTemplate.Spec.Labels == nil && envTemplate.Spec.ApplicationId == nil {
			return "", errors.New("애플리케이션 아이디 혹은 라벨이 입력되어야함")
		} else if envTemplate.Spec.Labels != nil {
			param := map[string]string{"labels": strings.Join(envTemplate.Spec.Labels, ",")}

			_, err := api.SendPut("/"+workspaceId+"/application/env", header, param, request)
			if err != nil {
				return "", err
			}
		} else if envTemplate.Spec.ApplicationId != nil {
			applicationId := *envTemplate.Spec.ApplicationId
			_, err := api.SendPut("/"+workspaceId+"/application/"+applicationId+"/env", header, map[string]string{}, request)
			if err != nil {
				return "", err
			}
		}

		return "", nil
	} else if resourceType == "GLOBAL_ENV" || resourceType == "GE" {
		var globalEnvTemplate template.GlobalEnvTemplate
		err := yaml.Unmarshal(content, &globalEnvTemplate)
		if err != nil {
			return "", err
		}

		request, err := json.Marshal(globalEnvTemplate.ToRequest())

		if err != nil {
			return "", err
		}

		_, err = api.SendPut("/workspace/"+globalEnvTemplate.Spec.WorkspaceId+"/env", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		return "", nil
	} else {
		return "", errors.New("지원되지 않는 리소스 타입입니다.")
	}
}
