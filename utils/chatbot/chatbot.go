package chatbot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Chatbot interface {
	HandleConnectionChatBot() echo.HandlerFunc
	ChatBot(input string) (string, error)
}

type ChatbotData struct {
	url    string
	client http.Client
}

var AiPayload = map[string]interface{}{
	"model": "gpt-3.5-turbo",
	"messages": []map[string]string{
		{"role": "system", "content": "Anda adalah seorang profesional pemilih sampah, kamu akan diberikan pertanyaan seputar sampah, kamu harus menjawab seputar jenis-jenis sampah dan penjelasan detail tentang sampah yang akan user input, jika user menginput selain sampah, anda akan menjawab 'anda tidak tahu' dan hanya menerima respon yang berhubungan dengan sampah saja"},
	},
}

func New() Chatbot {
	return &ChatbotData{
		url: "https://wgpt-production.up.railway.app/v1/chat/completions",
		client: http.Client{
			Timeout: time.Second * 60,
		},
	}
}

func (data *ChatbotData) HandleConnectionChatBot() echo.HandlerFunc {
	return func(c echo.Context) error {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			go func() {
				response, err := data.ChatBot(string(message))
				if err != nil {
					log.Println("Error processing message:", err)
					return
				}

				if response != "" {
					err = conn.WriteMessage(websocket.TextMessage, []byte(response))
					if err != nil {
						log.Println("Error writing message:", err)
						return
					}
				}
			}()
		}
		return nil
	}
}

func (data *ChatbotData) ChatBot(input string) (string, error) {
	// Jika pesan sistem belum ada, tambahkan ke payload
	if len(AiPayload["messages"].([]map[string]string)) == 1 {
		AiPayload["messages"] = append(AiPayload["messages"].([]map[string]string), map[string]string{"role": "system", "content": "Anda adalah seorang profesional pemilih sampah, kamu akan diberikan pertanyaan seputar sampah, kamu harus menjawab seputar jenis-jenis sampah dan penjelasan detail tentang sampah yang akan user input, jika user menginput selain sampah, anda akan menjawab 'anda tidak tahu' dan hanya menerima respon yang berhubungan dengan sampah saja"})
	}

	// Tambahkan pesan pengguna ke payload
	AiPayload["messages"] = append(AiPayload["messages"].([]map[string]string), map[string]string{"role": "user", "content": input})

	payload := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": []map[string]string{AiPayload["messages"].([]map[string]string)[len(AiPayload["messages"].([]map[string]string))-2], AiPayload["messages"].([]map[string]string)[len(AiPayload["messages"].([]map[string]string))-1]},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", data.url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := data.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// Tambahkan pesan asisten ke payload
	if len(result.Choices) > 0 {
		AiPayload["messages"] = append(AiPayload["messages"].([]map[string]string), map[string]string{"role": result.Choices[0].Message.Role, "content": result.Choices[0].Message.Content})
	}

	// Buat payload yang akan dikirim, hanya memasukkan pesan asisten terakhir
	// updatedPayload := map[string]interface{}{
	// 	"model":    "gpt-3.5-turbo",
	// 	"messages": []map[string]string{AiPayload["messages"].([]map[string]string)[len(AiPayload["messages"].([]map[string]string))-1]},
	// }

	// updatedJSONPayload, err := json.Marshal(updatedPayload)
	if err != nil {
		return "", err
	}

	return result.Choices[0].Message.Content, nil
}
