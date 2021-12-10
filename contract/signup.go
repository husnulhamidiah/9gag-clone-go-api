package contract

type SignupRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {}