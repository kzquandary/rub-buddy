package routes

import (
	"rub_buddy/configs"
	"rub_buddy/features/collectors"
	pickup "rub_buddy/features/pickup_request"
	"rub_buddy/features/users"
	bucket "rub_buddy/utils/bucket"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/users/login", uh.Login())
	e.POST("/users/register", uh.Register())
	e.GET("/users", uh.GetUser(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/users", uh.UpdateUser(), echojwt.JWT([]byte(cfg.Secret)))

}

func RouteCollector(e *echo.Echo, ch collectors.CollectorHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/collectors/login", ch.Login())
	e.POST("/collectors/register", ch.Register())
	e.GET("/collectors", ch.GetCollector(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/collectors", ch.UpdateCollector(), echojwt.JWT([]byte(cfg.Secret)))
}

func RoutePickup(e *echo.Echo, ph pickup.PickupRequestHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST("/pickup", ph.CreatePickupRequest(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/pickup", ph.GetAllPickupRequest(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/pickup/:id", ph.GetPickupRequestByID(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/pickup/:id", ph.DeletePickupRequestByID(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteMedia(e *echo.Echo, b bucket.BucketInterface) {
	e.POST("/media/upload", b.UploadFileHandler())
}

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
