package request

type SendAuthCodeRequest struct {
	Email string `json:"email"`
	Usage string `json:"usage"`
}

type CertificateAuthCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Usage string `json:"usage"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
