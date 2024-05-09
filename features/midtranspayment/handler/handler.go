package handler

import (
	"net/http"
	"rub_buddy/features/midtranspayment"
	"rub_buddy/helper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MidtransHandler struct {
	s midtranspayment.MidtransServiceInterface
	j helper.JWTInterface
}

func New(s midtranspayment.MidtransServiceInterface, j helper.JWTInterface) midtranspayment.MidtransHandlerInterface {
	return &MidtransHandler{
		s: s,
		j: j,
	}
}

func (h *MidtransHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		stringToken := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(stringToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		userData := h.j.ExtractToken(token)
		userId := userData["id"].(uint)
		var input = new(MidtransRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		var payment midtranspayment.Midtrans

		payment = midtranspayment.Midtrans{
			ID:     uuid.New().String(),
			Amount: input.Amount,
			UserID: userId,
		}

		res, err := h.s.GenerateSnapURL(payment)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error create transaction")
		}
		var response = new(MidtransResponse)
		response.ID = res.ID
		response.UserID = res.UserID
		response.Amount = res.Amount
		response.SnapURL = res.SnapURL
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "success create transaction", []interface{}{response}))
	}
}

func (h *MidtransHandler) VerifyPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		stringToken := c.Request().Header.Get("Authorization")
		_, err := h.j.ValidateToken(stringToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		orderId := c.Param("order_id")

		err = h.s.VerifyPayment(orderId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success verify payment", []interface{}{}))
	}
}
