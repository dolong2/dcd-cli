package template

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
