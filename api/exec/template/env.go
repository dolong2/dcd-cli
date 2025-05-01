package template

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
