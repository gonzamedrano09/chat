package entity

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Connection struct {
	ID         uuid.UUID
	Connection *websocket.Conn
}
