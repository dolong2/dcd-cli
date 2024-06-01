package exec

import (
	"encoding/json"
	"errors"
	"github.com/dolong2/dcd-cli/api"
	"os"
)

type metaData struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

type parsingMetaData struct {
	Metadata metaData `json:"metadata"`
}

type workspaceTemplate struct {
	Metadata metaData `json:"metadata"`
}

type workspaceRequest struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"title"`
	Description  string `json:"description"`
}

func CreateByPath(fileDirectory string) error {
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		return err
	}

	err = create(content)
	if err != nil {
		return err
	}
	return nil
}

func CreateByTemplate(rawTemplate string) error {
	err := create([]byte(rawTemplate))
	if err != nil {
		return err
	}

	return nil
}

func create(content []byte) error {
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

		_, err = api.SendPost("/workspace", header, request)
		if err != nil {
			return err
		}
	} else {
		return errors.New(" this resource type is not supported")
	}

	return nil
}
