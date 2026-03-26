package response

type CreateApplicationResponse struct {
	ApplicationId string `json:"applicationId"`
}

type ApplicationDetailResponse struct {
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	ApplicationType     string            `json:"applicationType"`
	GithubUrl           string            `json:"githubUrl"`
	Env                 map[string]string `json:"env"`
	Port                int               `json:"port"`
	ExternalPort        int               `json:"externalPort"`
	Version             string            `json:"version"`
	Status              string            `json:"status"`
	FailureReason       string            `json:"failureReason"`
	FailureReasonDetail string            `json:"failureReasonDetail"`
	InitialScripts      []string          `json:"initialScripts"`
	Labels              []string          `json:"labels"`
}

type ApplicationResponse struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ApplicationType string    `json:"applicationType"`
	GithubUrl       string    `json:"githubUrl"`
	Port            int       `json:"port"`
	ExternalPort    int       `json:"externalPort"`
	Version         string    `json:"version"`
	Status          string    `json:"status"`
	Labels          []string  `json:"labels"`
}

type ApplicationListResponse struct {
	Applications []ApplicationResponse `json:"list"`
}

type applicationSimpleResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
