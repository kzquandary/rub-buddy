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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}

		user, err := h.s.Login(input.Email, input.Password)

		if err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.UserNotFound, nil))
			}
			if strings.Contains(err.Error(), constant.EmailAndPasswordCannotBeEmpty) {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.EmailAndPasswordCannotBeEmpty, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}
		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}
		user, err := h.s.Register(users.User{
			Email:    input.Email,
			Password: HashedPassword,
			Name:     input.Name,
			Address:  input.Address,
			Gender:   input.Gender,
		})
		if err != nil {
			if strings.Contains(err.Error(), constant.EmailAlreadyExists) {
				return c.JSON(http.StatusConflict, helper.FormatResponse(false, constant.EmailAlreadyExists, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
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

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}
		token, err := h.jwt.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := h.jwt.ExtractToken(token)

		userDetails, err := h.s.GetUserByEmail(userData["email"].(string))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Get user success", []interface{}{userDetails}))
	}
}

func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		token, err := h.jwt.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := h.jwt.ExtractToken(token)

		var input = new(UpdateUserInput)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}

		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		var user = new(users.UserUpdate)
		user.ID = userData["id"].(uint)
		user.Email = input.Email
		user.Name = input.Name
		user.Address = input.Address
		user.Gender = input.Gender
		user.Password = HashedPassword

		err = h.s.UpdateUser(user)
		if err != nil {
			if strings.Contains(err.Error(), constant.EmailAlreadyExists) {
				return c.JSON(http.StatusConflict, helper.FormatResponse(false, constant.EmailAlreadyExists, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		var response = new(UserInfoResponse)
		response.ID = user.ID
		response.Email = user.Email
		response.Name = user.Name
		response.Address = user.Address
		response.Gender = user.Gender
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Update user success", []interface{}{response}))
	}
}
