package response

type VolumeListResponse struct {
	List []volumeSimpleResponse `json:"list"`
}

type volumeSimpleResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
