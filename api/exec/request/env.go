package request

type EnvPutRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Encryption bool   `json:"encryption"`
}

type EnvPutListRequest struct {
	EnvList []EnvPutRequest `json:"envList"`
}
