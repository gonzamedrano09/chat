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

	InsertConnectionToChatUser(connection *entity.Connection, chatUser *entity.ChatUser) error
	SelectOneConnectionOfChatUser(connection *entity.Connection, chatUser *entity.ChatUser, id uuid.UUID) error
	SelectAllConnectionsOfChatUser(connections *[]entity.Connection, chatUser *entity.ChatUser) error
	DeleteConnectionOfChatUser(chatUser *entity.ChatUser, id uuid.UUID) error
}
