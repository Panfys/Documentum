package models

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Client struct {
	Conn  *websocket.Conn
	Login string
	Mutex    sync.Mutex
	LastActive time.Time
}
