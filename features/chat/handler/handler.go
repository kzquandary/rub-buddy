package handler

import (
	"log"
	rubbuddychat "rub_buddy/features/chat"
	"rub_buddy/helper"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	c rubbuddychat.ChatServiceInterface
	j helper.JWTInterface
}

func NewHandler(c rubbuddychat.ChatServiceInterface, j helper.JWTInterface) *ChatHandler {
	return &ChatHandler{
		c: c,
		j: j,
	}
}

func (h *ChatHandler) GetChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		userData := h.j.ExtractToken(token)
		userId := userData["id"].(uint)
		userRole := userData["role"].(string)
		log.Print(userRole)
		log.Print(userId)
		query, err := h.c.GetChat(userId, userRole)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		var response []ChatInfo
		for _, chat := range query {
			chatInfo := ChatInfo{
				ID:                  chat.ID,
				PickupTransactionID: chat.PickupTransactionID,
			}
			userInfo := UserInfo{
				ID:   chat.UserID,
				Name: chat.UserName,
			}

			collectorInfo := CollectorInfo{
				ID:   chat.CollectorID,
				Name: chat.CollectorName,
			}
			chatInfo.User = userInfo
			chatInfo.Collector = collectorInfo
			response = append(response, chatInfo)
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{response}))
	}
}
