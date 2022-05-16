package todo

import "todo/entity"

type ICart interface {
	GetAllTodos() ([]entity.Todo, error)
	GetTodoById(id int) (entity.Todo, error)
	CreateTodo(book entity.Todo) error
	DeleteTodo(id, userID int) error
	UpdateTodo(id int, book entity.Todo) error
	CompleteTodo(todo entity.Todo) error
	ReOpen(todo entity.Todo) error
	GetAllTodosByIdUser(UserID int) ([]entity.Todo, error)
}
