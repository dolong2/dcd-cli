package template

import (
	"errors"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

type ApplicationTemplate struct {
	Metadata metaData                `json:"metadata" yaml:"metadata"`
	Spec     applicationSpecTemplate `json:"spec" yaml:"spec"`
}

func (template ApplicationTemplate) validateMetadata() error {
	if template.Metadata.Name == nil || template.Metadata.Description == nil {
		return errors.New("애플리케이션 메타데이터 정보가 올바르지 않습니다")
	}

	return nil
}

type applicationSpecTemplate struct {
	GithubUrl       string   `json:"githubUrl" yaml:"githubUrl"`
	ApplicationType string   `json:"applicationType" yaml:"applicationType"`
	Port            int      `json:"port" yaml:"port"`
	Version         string   `json:"version" yaml:"version"`
	Labels          []string `json:"labels" yaml:"labels"`
}

func (template ApplicationTemplate) ToRequest() (*request.ApplicationRequest, error) {
	err := template.validateMetadata()
	if err != nil {
		return nil, err
	}

	return &request.ApplicationRequest{
		Name:            *template.Metadata.Name,
		Description:     *template.Metadata.Description,
		GithubUrl:       template.Spec.GithubUrl,
		ApplicationType: template.Spec.ApplicationType,
		Port:            template.Spec.Port,
		Version:         template.Spec.Version,
		Labels:          template.Spec.Labels,
	}, nil
}
