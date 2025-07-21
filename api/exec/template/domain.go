package template

import "github.com/dolong2/dcd-cli/api/exec/request"

type DomainTemplate struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}

func (template DomainTemplate) ToRequest() request.CreateDomainRequest {
	return request.CreateDomainRequest{
		Name:        *template.Metadata.Name,
		Description: *template.Metadata.Description,
	}
}
