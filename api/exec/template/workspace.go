package template

import "errors"

type WorkspaceTemplate struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

func (template WorkspaceTemplate) ValidateMetadata() error {
	if template.Metadata.Name == nil || template.Metadata.Description == nil {
		return errors.New("워크스페이스 메타데이터 정보가 올바르지 않습니다")
	}

	return nil
}
