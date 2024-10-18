package requests

type MovieRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
}

//
//type UpdateMovieRequest struct {
//	Name        string `json:"name"`
//	Description string `json:"description"`
//	Genre       string `json:"genre"`
//}
