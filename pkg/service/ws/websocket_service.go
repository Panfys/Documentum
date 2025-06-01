package ws

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WSService interface {
	RegisterClient(client *models.Client)
	UnregisterClient(client *models.Client)
	HandleMessage(client *models.Client, message models.Message) error
	UpdateClientActivity(client *models.Client)
	SendToClient(client *models.Client, message models.Message) error
	Broadcast(message models.Message)
}

type WebSocketService struct {
	clients map[*models.Client]bool
	mu      sync.Mutex
	log     logger.Logger
}

func NewWebSocketService(log logger.Logger) *WebSocketService {
	ws := &WebSocketService{
		clients: make(map[*models.Client]bool),
		log:     log,
	}
	go ws.cleanupRoutine()
	return ws
}

func (s *WebSocketService) cleanupRoutine() {
	for {
		time.Sleep(30 * time.Second)
		s.mu.Lock()
		for client := range s.clients {
			if time.Since(client.LastActive) > 60*time.Second {
				client.Conn.Close()
				delete(s.clients, client)
				s.log.Info("отключение неактивного пользователя: " + client.Login)
			}
		}
		s.mu.Unlock()
	}
}

func (s *WebSocketService) RegisterClient(client *models.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for c := range s.clients {
		if c.Login == client.Login {
			if c.Agent != client.Agent || c.IP != client.IP {
				s.disconnectClient(c)
				//s.log.Info("дисконнект " + client.Login)
			}
			c.Mutex.Lock()
			c.Conn.Close()
			c.Conn = client.Conn
			c.LastActive = time.Now()
			c.Mutex.Unlock()
			//s.log.Info("обновление статуса клиента " + client.Login)
			return
		}
	}

	// Новый клиент
	client.LastActive = time.Now()
	s.clients[client] = true
	s.log.Info("подключение клиента " + client.Login)
}

func (s *WebSocketService) UnregisterClient(client *models.Client) {

	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	if time.Since(client.LastActive) > 30*time.Minute {
		client.Conn.Close()
		s.mu.Lock()
		delete(s.clients, client)
		s.mu.Unlock()
	}
}

func (s *WebSocketService) disconnectClient(client *models.Client) {
	var message models.Message
	message.Action = "disconnect"
	s.SendToClient(client, message) 
}

func (s *WebSocketService) UpdateClientActivity(client *models.Client) {
	client.Mutex.Lock()
	defer client.Mutex.Unlock()
	client.LastActive = time.Now()
}

func (s *WebSocketService) HandleMessage(client *models.Client, message models.Message) error {
	// Обработка PING-запроса
	if message.Action == "PING" {
		pongMessage := models.Message{
			Action: "PONG",
		}
	
		return s.SendToClient(client, pongMessage)
	}

	return nil
}

func (s *WebSocketService) SendToClient(client *models.Client, message models.Message) error {
	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return s.log.Error("Ошибка обработки в JSON при отправке PONG", err)
	}
	
	return client.Conn.WriteMessage(websocket.TextMessage, messageBytes)
}

func (s *WebSocketService) Broadcast(message models.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for client := range s.clients {
		if err := s.SendToClient(client, message); err != nil {
			s.log.Error("ошибка отправки сообщения всем пользователям, пользователю: ", client.Login, ", Error:", err)
		}
	}
}
