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
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		token, err := h.j.ValidateToken(tokenString)

		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractToken(token)

		if userData[constant.JWT_ROLE] != constant.RoleCollector {
			helper.UnauthorizedError(c)
		}

		collectorID := userData[constant.JWT_ID].(uint)

		var pickupTransactionInput PickupTransactionInput

		if err := c.Bind(&pickupTransactionInput); err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		var Transaction = pickuptransaction.PickupTransaction{
			CollectorID:     collectorID,
			PickupRequestID: pickupTransactionInput.PickupRequestID,
			TpsID:           pickupTransactionInput.TpsID,
		}
		pickupTransaction, err := h.p.CreatePickupTransaction(Transaction)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		var response = new(PickupTransactionCreate)
		response.ID = pickupTransaction.ID
		response.PickupRequestID = pickupTransaction.PickupRequestID
		response.PickupTime = pickupTransaction.PickupTime.Format(constant.ParseTime)
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.PickupTransactionCreateSuccess, []interface{}{response}))
	}
}

func (h *PickupTransactionHandler) GetAllPickupTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)

		token, err := h.j.ValidateToken(tokenString)

		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractToken(token)

		userId := userData[constant.JWT_ID].(uint)
		userRole := userData[constant.JWT_ROLE].(string)
		if userRole != constant.RoleCollector {
			helper.UnauthorizedError(c)
		}
		pickupTransaction, err := h.p.GetAllPickupTransaction(userId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		var response []PickupTransactionInfo
		for _, pt := range pickupTransaction {
			pickupTransactionInfo := PickupTransactionInfo{
				ID:         pt.ID,
				PickupTime: pt.PickupTime.Format(constant.ParseTime),
				TpsID:      pt.TpsID,
			}
			userInfo := UserInfo{
				ID:      pt.UserID,
				Name:    pt.UserName,
				Address: pt.UserAddress,
			}
			collectorInfo := CollectorInfo{
				ID:   pt.CollectorID,
				Name: pt.CollectorName,
			}
			pickupTransactionInfo.User = userInfo
			pickupTransactionInfo.Collector = collectorInfo
			response = append(response, pickupTransactionInfo)
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.PickupTransactionGetSuccess, []interface{}{response}))
	}
}

func (h *PickupTransactionHandler) GetPickupTransactionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		pickupTransactionID := c.Param("id")
		pickupTransactionIDInt, err := strconv.Atoi(pickupTransactionID)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(constant.ErrBadRequest), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		pickupTransaction, err := h.p.GetPickupTransactionByID(pickupTransactionIDInt)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(PickupTransactionInfo)
		response.ID = pickupTransaction.ID
		response.PickupTime = pickupTransaction.PickupTime.Format(constant.ParseTime)
		response.TpsID = pickupTransaction.TpsID
		response.User = UserInfo{
			ID:      pickupTransaction.UserID,
			Name:    pickupTransaction.UserName,
			Address: pickupTransaction.UserAddress,
		}
		response.Collector = CollectorInfo{
			ID:   pickupTransaction.CollectorID,
			Name: pickupTransaction.CollectorName,
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.PickupTransactionGetSuccess, []interface{}{response}))
	}
}
