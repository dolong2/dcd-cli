package response

type Profile struct {
	User          userProfile        `json:"user"`
	WorkspaceList []workspaceProfile `json:"workspaces"`
}

type userProfile struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type workspaceProfile struct {
	Id              string               `json:"id"`
	Title           string               `json:"title"`
	ApplicationList []applicationProfile `json:"applicationList"`
}

type applicationProfile struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
