package template

type metaData struct {
	ResourceType string  `json:"resourceType" yaml:"resourceType"`
	Name         *string `json:"name,omitempty" yaml:"name,omitempty"`
	Description  *string `json:"description,omitempty" yaml:"description,omitempty"`
}

type ParsingMetaData struct {
	Metadata metaData `json:"metadata" yaml:"metadata"`
}
