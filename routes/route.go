package routes

import (
	"rub_buddy/configs"
	"rub_buddy/constant/routesname"
	"rub_buddy/features/chat"
	"rub_buddy/features/collectors"
	"rub_buddy/features/midtranspayment"
	pickup "rub_buddy/features/pickup_request"
	pickuptransaction "rub_buddy/features/pickup_transaction"
	"rub_buddy/features/users"
	bucket "rub_buddy/utils/bucket"
	"rub_buddy/utils/chatbot"
	websocket "rub_buddy/utils/websocket"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST(routesname.UserLogin, uh.Login())
	e.POST(routesname.UserRegister, uh.Register())
	e.GET(routesname.UserBasePath, uh.GetUser(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT(routesname.UserBasePath, uh.UpdateUser(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteCollector(e *echo.Echo, ch collectors.CollectorHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST(routesname.CollectorLogin, ch.Login())
	e.POST(routesname.CollectorRegister, ch.Register())
	e.GET(routesname.CollectorBasePath, ch.GetCollector(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT(routesname.CollectorBasePath, ch.UpdateCollector(), echojwt.JWT([]byte(cfg.Secret)))
}

func RoutePickup(e *echo.Echo, ph pickup.PickupRequestHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST(routesname.PickupBasePath, ph.CreatePickupRequest(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET(routesname.PickupBasePath, ph.GetAllPickupRequest(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET(routesname.PickupById, ph.GetPickupRequestByID(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE(routesname.PickupById, ph.DeletePickupRequestByID(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteTransaction(e *echo.Echo, th pickuptransaction.PickupTransactionHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST(routesname.TransactionBasePath, th.CreatePickupTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET(routesname.TransactionBasePath, th.GetAllPickupTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET(routesname.TransactionById, th.GetPickupTransactionByID(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteChat(e *echo.Echo, ch chat.ChatHandlerInterface, cfg configs.ProgrammingConfig) {
	e.GET(routesname.ChatBasePath, ch.GetChat(), echojwt.JWT([]byte(cfg.Secret)))
}
func RouteWebsocket(e *echo.Echo, wh websocket.Websocket, cfg configs.ProgrammingConfig) {
	e.GET(routesname.ChatBasePath, wh.HandleConnection())
	e.POST(routesname.ChatMessage, wh.SendMessage(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteMedia(e *echo.Echo, b bucket.BucketInterface) {
	e.POST(routesname.MediaUpload, b.UploadFileHandler())
}

func RouteChatbot(e *echo.Echo, cb chatbot.Chatbot) {
	e.GET(routesname.ChatBot, cb.HandleConnectionChatBot())
}

func RouteMidtrans(e *echo.Echo, mh midtranspayment.MidtransHandlerInterface, cfg configs.ProgrammingConfig) {
	e.POST(routesname.PaymentBasePath, mh.CreateTransaction(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET(routesname.PaymentVerify, mh.VerifyPayment(), echojwt.JWT([]byte(cfg.Secret)))
}
