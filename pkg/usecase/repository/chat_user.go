package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/google/uuid"
)

type ChatUserRepositoryInterface interface {
	InsertChatUser(chatUser *entity.ChatUser) error
	SelectOne(chatUser *entity.ChatUser, id uint) error
	SelectAll(chatUsers *[]entity.ChatUser) error
	DeleteChatUser(id uint) error

	AddConnectionToChatUser(connection *entity.Connection, chatUser *entity.ChatUser)
	SelectOneConnectionOfChatUser(connection *entity.Connection, chatUser *entity.ChatUser, id uuid.UUID)
	SelectAllConnectionsOfChatUser(connections *[]entity.Connection, chatUser *entity.ChatUser)
	DeleteConnectionOfChatUser(chatUser *entity.ChatUser, id uuid.UUID)
}
