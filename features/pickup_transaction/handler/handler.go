package handler

import (
	"net/http"
	"rub_buddy/constant"
	pickuptransaction "rub_buddy/features/pickup_transaction"
	"rub_buddy/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PickupTransactionHandler struct {
	p pickuptransaction.PickupTransactionDataInterface
	j helper.JWTInterface
}

func NewHandler(data pickuptransaction.PickupTransactionDataInterface, jwt helper.JWTInterface) pickuptransaction.PickupTransactionHandlerInterface {
	return &PickupTransactionHandler{
		p: data,
		j: jwt,
	}
}

func (h *PickupTransactionHandler) CreatePickupTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := h.j.ExtractToken(token)

		if userData["role"] != "Collector" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		collectorID := userData["id"].(uint)

		var pickupTransactionInput PickupTransactionInput

		if err := c.Bind(&pickupTransactionInput); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		var Transaction = pickuptransaction.PickupTransaction{
			CollectorID:     collectorID,
			PickupRequestID: pickupTransactionInput.PickupRequestID,
			TpsID:           pickupTransactionInput.TpsID,
		}
		pickupTransaction, err := h.p.CreatePickupTransaction(Transaction)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Pickup Transaction Created", []interface{}{pickupTransaction}))
	}
}

func (h *PickupTransactionHandler) GetAllPickupTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		token, err := h.j.ValidateToken(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := h.j.ExtractToken(token)

		userId := userData["id"].(uint)
		pickupTransaction, err := h.p.GetAllPickupTransaction(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Pickup Transaction Retrieved", []interface{}{pickupTransaction}))
	}
}

func (h *PickupTransactionHandler) GetPickupTransactionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		pickupTransactionID := c.Param("id")
		pickupTransactionIDInt, err := strconv.Atoi(pickupTransactionID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}

		pickupTransaction, err := h.p.GetPickupTransactionByID(pickupTransactionIDInt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Pickup Transaction Retrieved", []interface{}{pickupTransaction}))
	}
}
