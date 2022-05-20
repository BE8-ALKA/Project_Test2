package todo

import (
	"fmt"
	"net/http"

	"todo/delivery/middlewares"
	view "todo/delivery/views"
	"todo/delivery/views/todo"
	"todo/delivery/views/user"
	"todo/entity"
	todoRepo "todo/repository/todo"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type TodoController struct {
	Repo  todoRepo.ITodo
	Valid *validator.Validate
}

func New(repo todoRepo.ITodo, valid *validator.Validate) *TodoController {
	return &TodoController{
		Repo:  repo,
		Valid: valid,
	}
}

func (tc *TodoController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpTodo todo.CreateRequest
		var resp todo.TodoResponse

		if err := c.Bind(&tmpTodo); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, todo.BadRequest())
		}

		if err := tc.Valid.Struct(tmpTodo); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusBadRequest, todo.BadRequest())
		}

		newTodo := entity.Todo{
			Todo:        tmpTodo.Todo,
			Description: tmpTodo.Description,
			Status:      tmpTodo.Status,
		}

		data, err := tc.Repo.CreateTodo(newTodo)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		resp = todo.TodoResponse{
			ID:          data.ID,
			Todo:        tmpTodo.Todo,
			Description: tmpTodo.Description,
			Status:      tmpTodo.Status,
		}

		log.Info("berhasil register")
		return c.JSON(http.StatusCreated, todo.RegisterSuccess(resp))
	}
}

func (tc *TodoController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		
		return c.JSON(http.StatusOK, ))
	}
}
