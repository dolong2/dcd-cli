package template

import (
	"errors"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

type DomainTemplate struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

func (template DomainTemplate) ToRequest() request.CreateDomainRequest {
	return request.CreateDomainRequest{
		Name:        *template.Metadata.Name,
		Description: *template.Metadata.Description,
	}
}

func (template DomainTemplate) ValidateMetadata() error {
	if template.Metadata.Name == nil || template.Metadata.Description == nil {
		return errors.New("도메인 메타데이터 정보가 올바르지 않습니다")
	}

	return nil
}
