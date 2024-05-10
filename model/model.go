package model

type MovieRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type MovieResponse struct {
	Message string `json:"message"`
}
