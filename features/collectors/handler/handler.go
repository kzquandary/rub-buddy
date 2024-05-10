package handler

import (
	"net/http"
	"rub_buddy/constant"
	"rub_buddy/features/collectors"
	"rub_buddy/helper"

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
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		collector, err := h.s.Register(collectors.Collectors{
			Email:    input.Email,
			Password: HashedPassword,
			Name:     input.Name,
			Gender:   input.Gender,
		})
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(RegisterResponse)
		response.ID = collector.ID
		response.Email = collector.Email
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.CollectorRegisterSuccess, []interface{}{response}))
	}
}

func (h *CollectorHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		err := c.Bind(input)

		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		collector, err := h.s.Login(input.Email, input.Password)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(LoginResponse)
		response.ID = collector.ID
		response.Email = collector.Email
		response.Token = collector.Token

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.CollectorLoginSuccess, []interface{}{response}))
	}
}

func (h *CollectorHandler) UpdateCollector() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		collectorsData := h.j.ExtractToken(token)

		var input = new(UpdateCollectorInput)

		err = c.Bind(input)
		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		HashedPassword, err := helper.HashPassword(input.Password)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var collector = new(collectors.CollectorUpdate)
		collector.ID = collectorsData[constant.JWT_ID].(uint)
		collector.Name = input.Name
		collector.Gender = input.Gender
		collector.Email = input.Email
		collector.Password = HashedPassword

		err = h.s.UpdateCollector(collector)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(CollectorInfoResponse)
		response.ID = collector.ID
		response.Email = collector.Email
		response.Name = collector.Name
		response.Gender = collector.Gender
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.CollectorUpdateSuccess, []interface{}{response}))
	}
}

func (h *CollectorHandler) GetCollector() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)

		token, err := h.j.ValidateToken(tokenString)

		if err != nil {
			helper.UnauthorizedError(c)
		}

		collectorData := h.j.ExtractToken(token)

		collectorDetails, err := h.s.GetCollectorByEmail(collectorData[constant.JWT_EMAIL].(string))

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(CollectorInfoResponse)
		response.ID = collectorDetails.ID
		response.Email = collectorDetails.Email
		response.Name = collectorDetails.Name
		response.Gender = collectorDetails.Gender
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.CollectorGetSuccess, []interface{}{response}))
	}
}
