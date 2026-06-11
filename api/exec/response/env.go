package response

type PutEnvResponse struct {
	EnvId string `json:"envId"`
}

type EnvListResponse struct {
	List []envSimpleResponse `json:"list"`
}

type envSimpleResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EnvResponse struct {
	Id          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Details     []envDetailResponse `json:"details"`
}

type envDetailResponse struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Encryption bool   `json:"encryption"`
}
