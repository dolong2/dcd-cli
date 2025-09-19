package response

type VolumeListResponse struct {
	List []volumeSimpleResponse `json:"list"`
}

type volumeSimpleResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type VolumeDetailResponse struct {
	Id          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	MountList   []VolumeMountResponse `json:"mountList"`
}

type VolumeMountResponse struct {
	MountPath       string                    `json:"mountPath"`
	ReadOnly        bool                      `json:"readOnly"`
	ApplicationInfo applicationSimpleResponse `json:"applicationInfo"`
}
