package handler

import (
	"documentum/pkg/logger"
	"documentum/pkg/models"
	"documentum/pkg/service/ws"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
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
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin:     func(r *http.Request) bool { return true },
		EnableCompression: true,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.log.Error("ошибка соединения ws:", err)
		return
	}

	login := r.Context().Value(models.LoginKey).(string)
	agent := r.Context().Value(models.UserAgentKey).(string)
	ip := r.Context().Value(models.IPKey).(string)

	client := &models.Client{
		Conn:  conn,
		Login: login,
		Agent: agent,
		IP: ip,
	}
	h.service.RegisterClient(client)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.log.Error("Соединение с "+client.Login+" закрыто. Ошибка:", err)
			break
		}

		// Обновляем время активности клиента
		h.service.UpdateClientActivity(client)

		// Декодируем JSON в структуру Message
		var message models.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			h.log.Error("Ошибка декодирования JSON:", err)
			continue // Пропускаем некорректное сообщение
		}

		// Обрабатываем сообщение
		if err := h.service.HandleMessage(client, message); err != nil {
			h.log.Error("Ошибка обработки сообщения:", err)
		}
	}
}
