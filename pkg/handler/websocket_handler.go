package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/ws"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	log     logger.Logger
	service ws.WSService
}

func NewWebSocketHandler(log logger.Logger, service ws.WSService) *WebSocketHandler {
	return &WebSocketHandler{
		log:     log,
		service: service,
	}
}

func (h *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.log.Error("ошибка соединения ws:", err)
		return
	}

	login := r.Context().Value(models.LoginKey).(string)

	client := &models.Client{
		Conn:   conn,
		Login: login,
	}
	h.service.RegisterClient(client)

	// Отправка подтверждения
	if err := conn.WriteMessage(websocket.TextMessage, []byte(login)); err != nil {
		h.log.Error("Failed to send OK:", err)
		return
	}

	// Обработка сообщений
	for {
		_, msg, err := conn.ReadMessage() 
		if err != nil {
			//h.log.Error("соединение с " + login, " закрыто, Error:", err)
			break
		}
		
		// Обновляем время активности
		h.service.UpdateClientActivity(client)
		
		if err := h.service.HandleMessage(client, msg); err != nil {
			h.log.Error("ошибка отправки сообщения: ", err)
		}
	}
}