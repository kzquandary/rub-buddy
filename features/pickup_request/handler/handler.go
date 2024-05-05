package handler

import (
	"log"
	"net/http"
	"rub_buddy/constant"
	pickuprequest "rub_buddy/features/pickup_request"
	"rub_buddy/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type PickupRequestHandler struct {
	s pickuprequest.PickupRequestServiceInterface
	j helper.JWTInterface
}

func NewHandler(s pickuprequest.PickupRequestServiceInterface, jwt helper.JWTInterface) *PickupRequestHandler {
	return &PickupRequestHandler{
		s: s,
		j: jwt,
	}
}

func (h *PickupRequestHandler) CreatePickupRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		pickupRequestData := h.j.ExtractToken(token)
		var input = new(PickupRequestInput)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}

		UserID := pickupRequestData["id"].(uint)
		UserAddress := pickupRequestData["address"].(string)
		Earnings := input.Weight * 3000
		pickupRequest := pickuprequest.PickupRequest{
			UserID:      UserID,
			Weight:      input.Weight,
			Address:     UserAddress,
			Description: input.Description,
			Earnings:    Earnings,
			Image:       input.Image,
		}

		_, err = h.s.CreatePickupRequest(pickupRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Pickup request created", nil))
	}
}

func (h *PickupRequestHandler) GetAllPickupRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		log.Print(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		pickupRequestData := h.j.ExtractToken(token)

		userRole := pickupRequestData["role"].(string)
		if userRole != "Collector" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		pickupRequests, err := h.s.GetAllPickupRequest()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Get all pickup request success", []interface{}{pickupRequests}))
	}
}

func (h *PickupRequestHandler) GetPickupRequestByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		pickupRequestID := c.Param("id")
		pickupRequestIDInt, err := strconv.Atoi(pickupRequestID)
		pickupRequest, err := h.s.GetPickupRequestByID(pickupRequestIDInt)
		if err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.NotFound, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Get pickup request success", []interface{}{pickupRequest}))
	}
}

func (h *PickupRequestHandler) DeletePickupRequestByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := h.j.ExtractToken(token)

		if userData["role"] != "User" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		UserID := userData["id"].(uint)

		pickupRequestID := c.Param("id")

		pickupRequestIDInt, err := strconv.Atoi(pickupRequestID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}
		err = h.s.DeletePickupRequestByID(pickupRequestIDInt, UserID)
		if err != nil {
			if strings.Contains(err.Error(), constant.NotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.NotFound, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Delete pickup request success", nil))
	}
}
