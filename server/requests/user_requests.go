package requests

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
