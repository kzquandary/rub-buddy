package handler

import (
	"net/http"
	"rub_buddy/constant"
	"rub_buddy/features/users"
	"rub_buddy/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s   users.UserServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(s users.UserServiceInterface, jwt helper.JWTInterface) *UserHandler {
	return &UserHandler{
		s:   s,
		jwt: jwt,
	}
}

func (h *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(input); err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		user, err := h.s.Login(input.Email, input.Password)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(LoginResponse)
		response.ID = user.ID
		response.Email = user.Email
		response.Token = user.Token
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.UserLoginSuccess, []interface{}{response}))

	}
}

func (h *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(input); err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}
		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user, err := h.s.Register(users.User{
			Email:    input.Email,
			Password: HashedPassword,
			Name:     input.Name,
			Address:  input.Address,
			Gender:   input.Gender,
		})
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		var response = new(RegisterResponse)
		response.ID = user.ID
		response.Email = user.Email
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.UserRegisterSuccess, []interface{}{response}))

	}
}

func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)

		if tokenString == "" {
			helper.UnauthorizedError(c)
		}
		token, err := h.jwt.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.jwt.ExtractToken(token)

		userDetails, err := h.s.GetUserByEmail(userData[constant.JWT_EMAIL].(string))
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(UserInfoResponse)
		response.ID = userDetails.ID
		response.Email = userDetails.Email
		response.Name = userDetails.Name
		response.Address = userDetails.Address
		response.Gender = userDetails.Gender
		response.Balance = userDetails.Balance
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.UserGetSuccess, []interface{}{response}))
	}
}

func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)

		token, err := h.jwt.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.jwt.ExtractToken(token)

		var input = new(UpdateUserInput)

		if err := c.Bind(input); err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var user = new(users.UserUpdate)
		user.ID = userData[constant.JWT_ID].(uint)
		user.Email = input.Email
		user.Name = input.Name
		user.Address = input.Address
		user.Gender = input.Gender
		user.Password = HashedPassword

		err = h.s.UpdateUser(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(UserInfoResponse)
		response.ID = user.ID
		response.Email = user.Email
		response.Name = user.Name
		response.Address = user.Address
		response.Gender = user.Gender
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.UserUpdateSuccess, []interface{}{response}))
	}
}
