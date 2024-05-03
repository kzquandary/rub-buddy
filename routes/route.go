package routes

import (
	"rub_buddy/configs"
	"rub_buddy/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())
	e.GET("/user", uh.GetUser(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/user", uh.UpdateUser(), echojwt.JWT([]byte(cfg.Secret)))

}

// func PickupRouter(e *echo.Echo, ph pickup.PickupHandlerInterface, cfg configs.ProgrammingConfig) {
// 	e.POST("/pickup", ph.Pickup(), echojwt.JWT([]byte(cfg.Secret)))
// 	e.GET("/pickup", ph.GetPickup(), echojwt.JWT([]byte(cfg.Secret)))
// 	e.GET("/pickup/{id}", ph.GetPickupByID(), echojwt.JWT([]byte(cfg.Secret)))
// 	e.DELETE("/pickup/{id}", ph.DeletePickup(), echojwt.JWT([]byte(cfg.Secret)))
// }

// func TransactionRouter(e *echo.Echo, th transaction.TransactionHandlerInterface, cfg configs.ProgrammingConfig) {
// 	e.POST("/transaction", th.Transaction(), echojwt.JWT([]byte(cfg.Secret)))
// 	e.GET("/transaction", th.GetTransaction(), echojwt.JWT([]byte(cfg.Secret)))
// 	e.GET("/transaction/{id}", th.GetTransactionByID(), echojwt.JWT([]byte(cfg.Secret)))
// }

// func ChatRouter(e *echo.Echo, ch chat.ChatHandlerInterface, cfg configs.ProgrammingConfig) {
// 	e.POST("/chat", ch.Chat(), echojwt.JWT([]byte(cfg.Secret)))
// 	e.GET("/chat/{id}", ch.GetChatByID(), echojwt.JWT([]byte(cfg.Secret)))
// }

// func UtilsRouter(e *echo.Echo, uh media.MediaHandlerInterface, cfg configs.ProgrammingConfig) {
// 	e.GET("/media/upload", uh.UploadMedia(), echojwt.JWT([]byte(cfg.Secret)))
// }

// func ChatBotRouter(e *echo.Echo, uh chatbot.ChatBotHandlerInterface, cfg configs.ProgrammingConfig) {
// 	e.POST("/chatbot", uh.ChatBot(), echojwt.JWT([]byte(cfg.Secret)))
// }
