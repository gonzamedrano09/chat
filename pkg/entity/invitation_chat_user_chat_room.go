package entity

import (
	"github.com/google/uuid"
)

type InvitationChatUserChatRoom struct {
	ID       uuid.UUID
	ChatUser *ChatUser
	ChatRoom *ChatRoom
}
