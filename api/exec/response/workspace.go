package response

type CreateWorkspaceResponse struct {
	WorkspaceId string `json:"workspaceId"`
}

type WorkspaceListResponse struct {
	List []WorkspaceResponse `json:"list"`
}

type WorkspaceResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type WorkspaceDetailResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
