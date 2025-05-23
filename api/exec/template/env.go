package template

import "github.com/dolong2/dcd-cli/api/exec/request"

type EnvTemplate struct {
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

func (template EnvTemplate) ToRequest() request.EnvPutListRequest {
	var envRequestList []request.EnvPutRequest

	envList := template.Spec.EnvList
	for _, env := range envList {
		envRequestList = append(envRequestList, request.EnvPutRequest{
			Key:        env.Key,
			Value:      env.Value,
			Encryption: env.Encryption,
		})
	}

	return request.EnvPutListRequest{EnvList: envRequestList}
}
