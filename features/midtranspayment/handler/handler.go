package handler

import (
	"net/http"
	"rub_buddy/constant"
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
		stringToken := c.Request().Header.Get(constant.HeaderAuthorization)
		token, err := h.j.ValidateToken(stringToken)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractToken(token)
		userId := userData[constant.JWT_ID].(uint)
		var input = new(MidtransRequest)
		if err := c.Bind(input); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, []interface{}{}))
		}

		payment := midtranspayment.Midtrans{
			ID:     uuid.New().String(),
			Amount: input.Amount,
			UserID: userId,
		}

		res, err := h.s.GenerateSnapURL(payment)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		var response = new(MidtransResponse)
		response.ID = res.ID
		response.UserID = res.UserID
		response.Amount = res.Amount
		response.SnapURL = res.SnapURL
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.PAYMENT_CREATED, []interface{}{response}))
	}
}

func (h *MidtransHandler) VerifyPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		stringToken := c.Request().Header.Get(constant.HeaderAuthorization)
		_, err := h.j.ValidateToken(stringToken)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		orderId := c.Param("order_id")

		err = h.s.VerifyPayment(orderId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.PAYMENT_VERIFIED, []interface{}{}))
	}
}
