package todo

import (
	"net/http"
)

type CreateResponse struct {
	Data  TodoResponse
	Token string
}

type TodoResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RegisterSuccess(data TodoResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil register Todo baru",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data Todo",
		"status":  false,
		"data":    nil,
	}
}
