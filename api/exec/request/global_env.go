package request

type GlobalEnvPutRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Encryption bool   `json:"encryption"`
}

type GlobalEnvPutListRequest struct {
	EnvList []GlobalEnvPutRequest `json:"envList"`
}
