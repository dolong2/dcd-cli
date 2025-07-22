package response

type CreateDomainResponse struct {
	DomainId string `json:"domainId"`
}

type DomainResponse struct {
	DomainId    string                     `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Application *applicationSimpleResponse `json:"application"`
}

type DomainListResponse struct {
	Domains []DomainResponse `json:"list"`
}
