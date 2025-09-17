package request

type VolumeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MountVolumeRequest struct {
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}
