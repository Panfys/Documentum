package models

import (
	
	"sync"
	"time"
	"github.com/gorilla/websocket"
	"encoding/json"
)

type Client struct {
	Conn  *websocket.Conn
	Login string
	Mutex    sync.Mutex
	LastActive time.Time
	Agent string
	IP string
}

type Message struct {
	Action string `json:"action"`
	Content json.RawMessage `json:"content,omitempty"`
}
type UpdDocFamConten struct {
	Type string `json:"type"`
	DocID string `json:"docID"`
	Familiar string `json:"familiar"` 
}

