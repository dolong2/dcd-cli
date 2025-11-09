package request

type ApplicationRequest struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	GithubUrl       string   `json:"githubUrl"`
	ApplicationType string   `json:"applicationType"`
	Port            int      `json:"port"`
	Version         string   `json:"version"`
	InitialScripts  []string `json:"initialScripts"`
	Labels          []string `json:"labels"`
}
