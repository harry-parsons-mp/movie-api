package requests

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type AuthRequest struct {
	Username string `json:"username"`
}

//type UpdateUserRequest struct {
//	Name     string `json:"name"`
//	Username string `json:"username"`
//}
