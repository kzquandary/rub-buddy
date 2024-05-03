package handler

import (
	"net/http"
	"rub_buddy/constant"
	"rub_buddy/features/users"
	"rub_buddy/helper"
	"strings"

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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to process request", err.Error()))
		}

		user, err := h.s.Login(input.Email, input.Password)

		if err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("User not found", err.Error()))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Internal server error", err.Error()))
		}

		var response = new(LoginResponse)
		response.ID = user.ID
		response.Email = user.Email
		response.Token = user.Token
		return c.JSON(http.StatusOK, helper.FormatResponse("Login success", response))

	}
}

func (h *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to process request", err.Error()))
		}
		user, err := h.s.Register(users.User{
			Email:    input.Email,
			Password: input.Password,
		})
		if err != nil {
			if strings.Contains(err.Error(), constant.EmailAlreadyExists) {
				return c.JSON(http.StatusConflict, helper.FormatResponse("Email already exists", err.Error()))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Internal server error", err.Error()))
		}
		var response = new(RegisterResponse)
		response.ID = user.ID
		response.Email = user.Email
		return c.JSON(http.StatusCreated, helper.FormatResponse("Register success", response))

	}
}

func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*users.User)
		if err := h.s.GetUser(user); err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("User not found", err.Error()))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Internal server error", err.Error()))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("Get user success", user))

	}
}

func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*users.User)
		if err := h.s.UpdateUser(user); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Internal server error", err.Error()))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("Update user success", user))
	}
}
