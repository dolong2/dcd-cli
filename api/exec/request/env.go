package request

type EnvPutRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Encryption bool   `json:"encryption"`
}

type EnvPutListRequest struct {
	Name                 string          `json:"name"`
	Description          string          `json:"description"`
	Details              []EnvPutRequest `json:"details"`
	ApplicationIdList    []string        `json:"applicationIdList"`
	ApplicationLabelList []string        `json:"applicationLabelList"`
}
