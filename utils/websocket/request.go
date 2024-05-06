package websocket

type MessageInput struct {
	ChatID  string `json:"chat_id"`
	Content string `json:"content"`
}
