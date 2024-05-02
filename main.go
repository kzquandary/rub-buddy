package main

import (
	"rub_buddy/config"
	controllers "rub_buddy/controllers/user"
	"rub_buddy/drivers/postgres"
	"rub_buddy/drivers/postgres/user"
	"rub_buddy/routes"
	"rub_buddy/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitConfigMySQL()
	db := postgres.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

	userRepo := user.NewUserRepo(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	routes := routes.RouteController{
		UserController: userController,
	}

	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
