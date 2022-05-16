package routes

import (
	"net/http"
	"todo/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc user.IUserController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	apiGroup := e.Group("/api/v1")

	// user login and register enpoints route
	apiGroup.POST("/login", uc.Login())
	apiGroup.POST("/register", uc.Register())

}
