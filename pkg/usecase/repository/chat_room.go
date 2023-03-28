package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/google/uuid"
)

type ChatRoomRepositoryInterface interface {
	InsertChatRoom(chatRoom *entity.ChatRoom) error
	SelectOne(chatRoom *entity.ChatRoom, id uuid.UUID) error
	SelectAll(chatRooms *[]entity.ChatRoom) error
	DeleteChatRoom(id uuid.UUID) error

	AddChatUserToChatRoom(chatUser *entity.ChatUser, chatRoom *entity.ChatRoom)
	SelectOneChatUserOfChatRoom(chatUser *entity.ChatUser, chatRoom *entity.ChatRoom, id uint)
	SelectAllChatUsersOfChatRoom(chatUsers *[]entity.ChatUser, chatRoom *entity.ChatRoom)
	DeleteChatUserOfChatRoom(chatRoom *entity.ChatUser, id uint)
}
