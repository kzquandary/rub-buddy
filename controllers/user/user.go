package controllers

import (
	"net/http"
	"rub_buddy/controllers/base"
	"rub_buddy/controllers/user/request"
	"rub_buddy/controllers/user/response"
	"rub_buddy/entities"
	"rub_buddy/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase entities.UseCaseInterface
}

func (uc *UserController) Register(c echo.Context) error {
	var userRegister request.UserRegister
	c.Bind(&userRegister)

	user, err := uc.userUseCase.Register(userRegister.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	userResponse := response.FromUseCaseToRegister(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", userResponse))
}

//	func (uc *UserController) Login(c echo.Context) error {
//		var userLogin request.UserLogin
//		c.Bind(&userLogin)
//		user, err := uc.userUseCase.Login(userLogin.ToEntities())
//		if err != nil {
//			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
//		}
//		userResponse := response.FromUseCase(&user)
//		return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
//	}
func NewUserController(userUseCase entities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}
