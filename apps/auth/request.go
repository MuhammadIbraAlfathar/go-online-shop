package auth

type RequestRegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
