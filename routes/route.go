package routes

import (
	controllers "rub_buddy/controllers/user"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	UserController *controllers.UserController
	
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/register", r.UserController.Register)
}
