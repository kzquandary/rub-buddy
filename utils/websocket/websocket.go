package websocket

import (
	"log"
	"net/http"
	"rub_buddy/constant"
	"rub_buddy/features/chat"
	chatmessage "rub_buddy/features/chat_message"
	"rub_buddy/helper"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Message struct {
	Sender    uint      `json:"sender"`
	Receiver  uint      `json:"receiver"`
	Content   string    `json:"content"`
	ChatID    uint      `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	clients      = make(map[*websocket.Conn]struct{})
	addClient    = make(chan *websocket.Conn)
	removeClient = make(chan *websocket.Conn)
	messages     = make(chan Message)
)

type SendMessageRequest struct {
	ChatID  uint   `json:"chat_id"`
	Content string `json:"content"`
}

type Websocket interface {
	HandleConnection() echo.HandlerFunc
	isChatIDExists(chatID uint) bool
	loadMessagesFromDB(chatID uint, conn *websocket.Conn) error
	saveMessageToDB(msg Message) error
	SendMessage() echo.HandlerFunc
}

type WebsocketData struct {
	db *gorm.DB
	j  helper.JWTInterface
}

func New(db *gorm.DB, j helper.JWTInterface) Websocket {
	return &WebsocketData{
		db: db,
		j:  j,
	}
}

func (data *WebsocketData) HandleConnection() echo.HandlerFunc {
	return func(c echo.Context) error {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		chatID := c.QueryParam("chat_id")
		tokenString := c.Request().Header.Get("Authorization")

		if chatID == "" || tokenString == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}

		token, err := data.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := data.j.ExtractToken(token)
		userID := userData["id"].(uint)
		chatIDInt, err := strconv.Atoi(chatID)
		chatIDUint := uint(chatIDInt)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.BadRequest, nil))
		}

		var chatInfo struct {
			ChatID      uint `json:"chat_id"`
			UserID      uint `json:"user_id"`
			CollectorID uint `json:"collector_id"`
		}
		query := `SELECT c.id AS chat_id,
			pr.user_id,
			p.collector_id
			FROM chats c
			JOIN pickup_transactions p ON c.pickup_transaction_id = p.id
			JOIN pickup_requests pr ON p.pickup_request_id = pr.id
			WHERE c.id = ?`
		result := data.db.Raw(query, chatIDUint).Scan(&chatInfo)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "Chat ID not found", nil))
		}

		var sender, receiver uint

		if userID == chatInfo.UserID {
			sender = userID
			receiver = chatInfo.CollectorID
		} else if userID == chatInfo.CollectorID {
			sender = userID
			receiver = chatInfo.UserID
		} else {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		defer conn.Close()

		err = data.loadMessagesFromDB(chatIDUint, conn)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, constant.InternalServerError, nil))
		}

		addClient <- conn

		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				removeClient <- conn
				break
			}
			msg.Sender = sender
			msg.Receiver = receiver
			msg.ChatID = chatIDUint

			messages <- msg
		}
		return nil
	}
}

func (data *WebsocketData) isChatIDExists(chatID uint) bool {
	var count int64
	data.db.Model(&chat.Chat{}).Where("id = ?", chatID).Count(&count)
	return count > 0
}

func (data *WebsocketData) loadMessagesFromDB(chatID uint, conn *websocket.Conn) error {

	var chatMessages []chatmessage.ChatMessage
	result := data.db.Table("chat_messages").Where("chat_id = ?", chatID).Order("created_at ASC").Find(&chatMessages)
	if result.Error != nil {
		return result.Error
	}

	// Kirim pesan ke klien
	for _, chatMsg := range chatMessages {
		msg := Message{
			Sender:    chatMsg.Sender,
			Receiver:  chatMsg.Receiver,
			Content:   chatMsg.Content,
			ChatID:    chatMsg.ChatID,
			CreatedAt: chatMsg.CreatedAt,
		}
		err := conn.WriteJSON(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (data *WebsocketData) HandleMessages() {
	for {
		select {
		case msg := <-messages:
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Println("Write error:", err)
					client.Close()
					delete(clients, client)
				}
			}

			err := data.saveMessageToDB(msg)
			if err != nil {
				log.Println("Database error:", err)
			}
		case client := <-addClient:
			clients[client] = struct{}{}
		case client := <-removeClient:
			delete(clients, client)
		}
	}
}

func (data *WebsocketData) saveMessageToDB(msg Message) error {
	var chat chat.Chat
	result := data.db.First(&chat, msg.ChatID)
	if result.Error != nil {
		return result.Error
	}

	// Buat pesan chat baru
	chatMessage := chatmessage.ChatMessage{
		ChatID:   msg.ChatID,
		Sender:   msg.Sender,
		Receiver: msg.Receiver,
		Content:  msg.Content,
	}

	// Simpan pesan chat ke database
	result = data.db.Create(&chatMessage)
	return result.Error
}

func (data *WebsocketData) SendMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse JSON dari body request
		tokenString := c.Request().Header.Get("Authorization")
		token, err := data.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		userData := data.j.ExtractToken(token)
		userID := userData["id"].(uint)

		// Ambil chat_id dan content dari body request
		var msg Message
		if err := c.Bind(&msg); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Invalid JSON", nil))
		}

		// Query untuk mendapatkan informasi chat
		var chatInfo struct {
			ChatID          uint `json:"chat_id"`
			PickupRequestID uint `json:"pickup_request_id"`
			UserID          uint `json:"user_id"`
			CollectorID     uint `json:"collector_id"`
		}
		query := `SELECT c.id AS chat_id,
			p.pickup_request_id,
			pr.user_id,
			p.collector_id
			FROM chats c
			JOIN pickup_transactions p ON c.pickup_transaction_id = p.id
			JOIN pickup_requests pr ON p.pickup_request_id = pr.id
			WHERE c.id = ?`
		result := data.db.Raw(query, msg.ChatID).Scan(&chatInfo)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "Chat ID not found", nil))
		}

		var sender, receiver uint
		if userID == chatInfo.UserID {
			sender = chatInfo.UserID
			receiver = chatInfo.CollectorID
		} else if userID == chatInfo.CollectorID {
			sender = chatInfo.CollectorID
			receiver = chatInfo.UserID
		} else {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, "Unauthorized", nil))
		}

		// Isi informasi pesan
		msg.Sender = sender
		msg.Receiver = receiver
		msg.CreatedAt = time.Now()

		// Kirim pesan ke channel messages
		messages <- msg

		// Memberikan respons bahwa pesan telah berhasil dikirim
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Message sent", nil))
	}
}
