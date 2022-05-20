package todo

import "github.com/labstack/echo/v4"

type ITodoController interface {
	Create() echo.HandlerFunc
	Show() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
