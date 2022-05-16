package main

import (
	"fmt"
	"log"
	"todo/config"
	"todo/delivery/routes"
	"todo/entity"

	userController "todo/delivery/controllers/user"
	userRepo "todo/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	db.AutoMigrate(
		&entity.User{})

	e := echo.New()

	repoUser := userRepo.New(db)
	controllerUser := userController.New(repoUser, validator.New())
	routes.RegisterPath(e, controllerUser)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))

}
