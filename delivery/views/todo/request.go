package todo

type CreateRequest struct {
	Todo        string `json:"email"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateRequest struct {
	Todo        string `json:"email"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
