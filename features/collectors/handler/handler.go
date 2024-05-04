package handler

import (
	"net/http"
	"rub_buddy/constant"
	"rub_buddy/features/collectors"
	"rub_buddy/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type CollectorHandler struct {
	s collectors.CollectorServiceInterface
	j helper.JWTInterface
}

func NewHandler(s collectors.CollectorServiceInterface, jwt helper.JWTInterface) *CollectorHandler {
	return &CollectorHandler{
		s: s,
		j: jwt,
	}
}

func (h *CollectorHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		err := c.Bind(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		collector, err := h.s.Register(collectors.Collectors{
			Email:    input.Email,
			Password: HashedPassword,
			Name:     input.Name,
			Gender:   input.Gender,
		})
		if err != nil {
			if strings.Contains(err.Error(), constant.EmailAlreadyExists) {
				return c.JSON(http.StatusConflict, helper.FormatResponse(false, constant.EmailAlreadyExists, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		var response = new(RegisterResponse)
		response.ID = collector.ID
		response.Email = collector.Email
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Register success", []interface{}{response}))
	}
}

func (h *CollectorHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		err := c.Bind(input)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		collector, err := h.s.Login(input.Email, input.Password)

		if err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.UserNotFound, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		var response = new(LoginResponse)
		response.ID = collector.ID
		response.Email = collector.Email
		response.Token = collector.Token

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Login success", []interface{}{response}))
	}
}

func (h *CollectorHandler) UpdateCollector() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}
		collectorsData := h.j.ExtractToken(token)

		var input = new(UpdateCollectorInput)

		err = c.Bind(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		HashedPassword, err := helper.HashPassword(input.Password)
		var collector = new(collectors.Collectors)
		collector.ID = collectorsData["id"].(uint)
		collector.Name = input.Name
		collector.Gender = input.Gender
		collector.Email = input.Email
		collector.Password = HashedPassword

		err = h.s.UpdateCollector(collector)
		if err != nil {
			if strings.Contains(err.Error(), constant.EmailAlreadyExists) {
				return c.JSON(http.StatusConflict, helper.FormatResponse(false, constant.EmailAlreadyExists, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		var response = new(CollectorInfoResponse)
		response.ID = collector.ID
		response.Email = collector.Email
		response.Name = collector.Name
		response.Gender = collector.Gender
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Update collector success", []interface{}{response}))
	}
}

func (h *CollectorHandler) GetCollector() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		token, err := h.j.ValidateToken(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		collectorData := h.j.ExtractToken(token)

		collectorDetails, err := h.s.GetCollectorByEmail(collectorData["email"].(string))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Get collector success", []interface{}{collectorDetails}))
	}
}
