package todo

import "todo/entity"

type ITodo interface {
	GetAllTodos() ([]entity.Todo, error)
	GetTodoById(id int) (entity.Todo, error)
	CreateTodo(newTodo entity.Todo) (entity.Todo, error)
	DeleteTodo(id, userID int) error
	UpdateTodo(id int, todo entity.Todo) error
	CompleteTodo(todo entity.Todo) error
	ReOpen(todo entity.Todo) error
	GetAllTodosByIdUser(UserID int) ([]entity.Todo, error)
}
