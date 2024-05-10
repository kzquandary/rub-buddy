package handler

import (
	"net/http"
	"rub_buddy/constant"
	pickuprequest "rub_buddy/features/pickup_request"
	"rub_buddy/helper"
	"strconv"

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
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		pickupRequestData := h.j.ExtractToken(token)
		var input = new(PickupRequestInput)
		if err := c.Bind(input); err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, []interface{}{}))
		}

		UserID := pickupRequestData[constant.JWT_ID].(uint)
		UserAddress := pickupRequestData[constant.JWT_ADDRESS].(string)
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
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.PickupRequestCreateSuccess, []interface{}{}))
	}
}

func (h *PickupRequestHandler) GetAllPickupRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		pickupRequestData := h.j.ExtractToken(token)
		userRole := pickupRequestData[constant.JWT_ROLE].(string)
		if userRole != constant.RoleCollector {
			helper.UnauthorizedError(c)
		}

		pickupRequests, err := h.s.GetAllPickupRequest()
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		var response []PickupRequestInfo
		for _, pr := range pickupRequests {
			pickupRequestInfo := PickupRequestInfo{
				ID:          pr.ID,
				Weight:      pr.Weight,
				Address:     pr.Address,
				Description: pr.Description,
				Earnings:    pr.Earnings,
				Image:       pr.Image,
			}
			userInfo := UserInfo{
				ID:   pr.UserID,
				Name: pr.UserName, 
			}
			pickupRequestInfo.User = userInfo

			response = append(response, pickupRequestInfo)
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.PickupRequestGetSuccess, []interface{}{response}))
	}
}

func (h *PickupRequestHandler) GetPickupRequestByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		pickupRequestID := c.Param("id")
		pickupRequestIDInt, err := strconv.Atoi(pickupRequestID)
		pickupRequest, err := h.s.GetPickupRequestByID(pickupRequestIDInt)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}

		var response = new(PickupRequestInfo)
		response.ID = pickupRequest.ID
		response.Weight = pickupRequest.Weight
		response.Address = pickupRequest.Address
		response.Description = pickupRequest.Description
		response.Earnings = pickupRequest.Earnings
		response.Image = pickupRequest.Image
		userInfo := UserInfo{
			ID:   pickupRequest.UserID,
			Name: pickupRequest.UserName,
		}
		response.User = userInfo

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.PickupRequestGetSuccess, []interface{}{response}))
	}
}

func (h *PickupRequestHandler) DeletePickupRequestByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractToken(token)

		if userData[constant.JWT_ROLE].(string) != constant.RoleUser {
			helper.UnauthorizedError(c)
		}

		UserID := userData[constant.JWT_ID].(uint)

		pickupRequestID := c.Param("id")

		pickupRequestIDInt, err := strconv.Atoi(pickupRequestID)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(constant.ErrBadRequest), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		
		err = h.s.DeletePickupRequestByID(pickupRequestIDInt, UserID)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.PickupRequestDeleteSuccess, []interface{}{}))
	}
}
