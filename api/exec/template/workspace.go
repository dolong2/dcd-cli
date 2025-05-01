package template

import (
	"errors"
	"github.com/dolong2/dcd-cli/api/exec/request"
)

type WorkspaceTemplate struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

func (template WorkspaceTemplate) ValidateMetadata() error {
	if template.Metadata.Name == nil || template.Metadata.Description == nil {
		return errors.New("워크스페이스 메타데이터 정보가 올바르지 않습니다")
	}

	return nil
}

func (template WorkspaceTemplate) ToRequest() request.WorkspaceRequest {
	return request.WorkspaceRequest{
		Name:        *template.Metadata.Name,
		Description: *template.Metadata.Description,
	}
}
