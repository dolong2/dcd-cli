package template

import "github.com/dolong2/dcd-cli/api/exec/request"

type GlobalEnvTemplate struct {
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

func (template GlobalEnvTemplate) ToRequest() request.GlobalEnvPutListRequest {
	var envRequestList []request.GlobalEnvPutRequest

	envList := template.Spec.EnvList
	for _, env := range envList {
		envRequestList = append(envRequestList, request.GlobalEnvPutRequest{
			Key:        env.Key,
			Value:      env.Value,
			Encryption: env.Encryption,
		})
	}

	return request.GlobalEnvPutListRequest{EnvList: envRequestList}
}
