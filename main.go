package main

import (
	"fmt"
	"log"
	"net/http"
	cfg "rub_buddy/configs"
	"rub_buddy/helper"
	"rub_buddy/routes"
	"rub_buddy/utils/database"

	dataUser "rub_buddy/features/users/data"
	handlerUser "rub_buddy/features/users/handler"
	serviceUser "rub_buddy/features/users/service"

	"github.com/labstack/echo/v4"
)

func main() {
	var config = cfg.InitConfig()
	db, err := database.InitDB(*config)
	database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	jwtInterface := helper.New(config.Secret)

	userModel := dataUser.New(db)
	userService := serviceUser.New(userModel, jwtInterface)
	userController := handlerUser.NewHandler(userService, jwtInterface)

	routes.RouteUser(e, userController, *config)

	e.Logger.Debug(db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
