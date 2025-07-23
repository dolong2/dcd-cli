package request

type SetDomainRequest struct {
	Domain string `json:"domain"`
}

type CreateDomainRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ConnectDomainRequest struct {
	ApplicationId string `json:"applicationId"`
}
