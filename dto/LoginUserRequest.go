package dto

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
