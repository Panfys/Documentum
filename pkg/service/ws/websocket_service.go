package ws

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type WSService interface {
	RegisterClient(client *models.Client)
	UnregisterClient(client *models.Client)
	HandleMessage(client *models.Client, message []byte) error
	UpdateClientActivity(client *models.Client)
	SendToClient(client *models.Client, message []byte) error
	Broadcast(message []byte)
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
			if time.Since(client.LastActive) > 45*time.Second {
				client.Conn.Close()
				delete(s.clients, client)
				s.log.Info("удаление неактивного пользователя: " + client.Login)
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
	s.log.Info("регистрация клиента " + client.Login)
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

func (s *WebSocketService) UpdateClientActivity(client *models.Client) {
	client.Mutex.Lock()
	defer client.Mutex.Unlock()
	client.LastActive = time.Now()
}

func (s *WebSocketService) HandleMessage(client *models.Client, message []byte) error {

	if string(message) == "PING" {
		return s.SendToClient(client, []byte("PONG"))
	} else {
		s.log.Info("Получение сообщения от пользователя " + client.Login + ", Content: " + string(message))
	}

	return nil
}

func (s *WebSocketService) SendToClient(client *models.Client, message []byte) error {
	client.Mutex.Lock()
	defer client.Mutex.Unlock()
	if string(message) != "PONG" {
		s.log.Info("Отправлено сообщение пользователю " + client.Login + ", Content: " + string(message)) 
	}

	return client.Conn.WriteMessage(websocket.TextMessage, message)
}

func (s *WebSocketService) Broadcast(message []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for client := range s.clients {
		if err := s.SendToClient(client, message); err != nil {
			s.log.Error("ошибка отправки сообщения всем пользователям, пользователю: ", client.Login, ", Error:", err)
		}
	}
}
