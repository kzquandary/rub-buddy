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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Failed to process request", nil))
		}

		user, err := h.s.Login(input.Email, input.Password)

		if err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "User not found", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Internal server error", nil))
		}

		var response = new(LoginResponse)
		response.ID = user.ID
		response.Email = user.Email
		response.Token = user.Token
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Login success", []interface{}{response}))

	}
}

func (h *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Failed to process request", nil))
		}
		user, err := h.s.Register(users.User{
			Email:    input.Email,
			Password: input.Password,
			Name:     input.Name,
			Address:  input.Address,
			Gender:   input.Gender,
		})
		if err != nil {
			if strings.Contains(err.Error(), constant.EmailAlreadyExists) {
				return c.JSON(http.StatusConflict, helper.FormatResponse(false, "Email already exists", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Internal server error", nil))
		}
		var response = new(RegisterResponse)
		response.ID = user.ID
		response.Email = user.Email
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Register success", []interface{}{response}))

	}
}

func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		c.Logger().Info("Token: ", tokenString)
		token, err := h.jwt.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}
		userData := h.jwt.ExtractToken(token)

		userDetails, err := h.s.GetUserByEmail(userData["email"].(string))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Internal server error", nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Get user success", []interface{}{userDetails}))
	}
}

func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*users.User)
		if err := h.s.UpdateUser(user); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Internal server error", nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Update user success", []interface{}{user}))
	}
}
