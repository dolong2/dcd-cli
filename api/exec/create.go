package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"github.com/dolong2/dcd-cli/api/exec/response"
	"github.com/dolong2/dcd-cli/api/exec/template"
	"os"
)

func CreateByPath(fileDirectory string) error {
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	unmarshal, err := resolveFileExtension(fileDirectory)
	resourceId, err := create(content, unmarshal)
	if err != nil {
		return err
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
	_, err := create([]byte(rawTemplate), json.Unmarshal)
	if err != nil {
		return err
	}

	return nil
}

func create(content []byte, unmarshal func([]byte, interface{}) (err error)) (string, error) {
	var data template.ParsingMetaData
	err := unmarshal(content, &data)
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

	switch resourceType {
	case "WORKSPACE":
		var workspace template.WorkspaceTemplate
		err = unmarshal(content, &workspace)
		if err != nil {
			return "", err
		}

		createWorkspaceRequest, err := workspace.ToRequest()
		if err != nil {
			return "", err
		}
		request, err := json.Marshal(createWorkspaceRequest)
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
	case "APPLICATION":
		var application template.ApplicationTemplate
		err := unmarshal(content, &application)
		if err != nil {
			return "", err
		}

		createApplicationRequest, err := application.ToRequest()
		if err != nil {
			return "", err
		}
		request, err := json.Marshal(createApplicationRequest)
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
	case "ENV":
		var envTemplate template.EnvTemplate
		err := unmarshal(content, &envTemplate)
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

		if envTemplate.Spec.ApplicationLabelList == nil && envTemplate.Spec.ApplicationIdList == nil {
			return "", errors.New("애플리케이션 아이디 혹은 라벨이 입력되어야함")
		} else {
			_, err := api.SendPost("/"+workspaceId+"/env", header, map[string]string{}, request)
			if err != nil {
				return "", err
			}
		}

		return "", nil
	case "DOMAIN":
		var domainTemplate template.DomainTemplate
		err := unmarshal(content, &domainTemplate)
		if err != nil {
			return "", err
		}

		createDomainRequest, err := domainTemplate.ToRequest()
		if err != nil {
			return "", err
		}
		request, err := json.Marshal(createDomainRequest)
		if err != nil {
			return "", err
		}

		workspaceId, err := getWorkspaceId()
		if err != nil {
			return "", err
		}

		result, err := api.SendPost("/"+workspaceId+"/domain", header, map[string]string{}, request)
		if err != nil {
			return "", err
		}

		createDomainResponse := response.CreateDomainResponse{}
		err = json.Unmarshal(result, &createDomainResponse)
		if err != nil {
			return "", err
		}

		return createDomainResponse.DomainId, nil
	default:
		return "", errors.New("지원되지 않는 리소스 타입입니다")
	}
}
