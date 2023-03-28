package entity

import (
	"github.com/google/uuid"
)

type ChatRoom struct {
	ID          uuid.UUID
	Members     []*ChatUser
	Invitations []*ChatUser
}
