package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
	"github.com/google/uuid"
)

type ChatUserRepository struct {
	ChatUsers []entity.ChatUser
}

func NewChatUserRepository() repository.ChatUserRepositoryInterface {
	return &ChatUserRepository{}
}

func (cur *ChatUserRepository) InsertChatUser(chatUser *entity.ChatUser) error {
	cur.ChatUsers = append(cur.ChatUsers, *chatUser)
	return nil
}

func (cur *ChatUserRepository) SelectOne(chatUser *entity.ChatUser, id uint) error {
	for _, iteratedChatUser := range cur.ChatUsers {
		if iteratedChatUser.User.ID == id {
			chatUser = &iteratedChatUser
		}
	}
	return nil
}

func (cur *ChatUserRepository) SelectAll(chatUsers *[]entity.ChatUser) error {
	chatUsers = &cur.ChatUsers
	return nil
}

func (cur *ChatUserRepository) DeleteChatUser(id uint) error {
	for indexIteratedChatUser, iteratedChatUser := range cur.ChatUsers {
		if iteratedChatUser.User.ID == id {
			cur.ChatUsers = append(cur.ChatUsers[:indexIteratedChatUser], cur.ChatUsers[indexIteratedChatUser+1])
		}
	}
	return nil
}

func (cur *ChatUserRepository) InsertConnectionToChatUser(connection *entity.Connection, chatUser *entity.ChatUser) error {
	var chatUserSelected *entity.ChatUser
	for _, iteratedChatUser := range cur.ChatUsers {
		if iteratedChatUser.User.ID == chatUser.User.ID {
			chatUserSelected = &iteratedChatUser
		}
	}
	chatUserSelected.Connections = append(chatUserSelected.Connections, *connection)
	return nil
}

func (cur *ChatUserRepository) SelectOneConnectionOfChatUser(connection *entity.Connection, chatUser *entity.ChatUser, id uuid.UUID) error {
	var chatUserSelected *entity.ChatUser
	for _, iteratedChatUser := range cur.ChatUsers {
		if iteratedChatUser.User.ID == chatUser.User.ID {
			chatUserSelected = &iteratedChatUser
		}
	}
	for _, iteratedConnection := range chatUserSelected.Connections {
		if iteratedConnection.ID == id {
			connection = &iteratedConnection
		}
	}
	return nil
}

func (cur *ChatUserRepository) SelectAllConnectionsOfChatUser(connections *[]entity.Connection, chatUser *entity.ChatUser) error {
	var chatUserSelected *entity.ChatUser
	for _, iteratedChatUser := range cur.ChatUsers {
		if iteratedChatUser.User.ID == chatUser.User.ID {
			chatUserSelected = &iteratedChatUser
		}
	}
	connections = &chatUserSelected.Connections
	return nil
}

func (cur *ChatUserRepository) DeleteConnectionOfChatUser(chatUser *entity.ChatUser, id uuid.UUID) error {
	var chatUserSelected *entity.ChatUser
	for _, iteratedChatUser := range cur.ChatUsers {
		if iteratedChatUser.User.ID == chatUser.User.ID {
			chatUserSelected = &iteratedChatUser
		}
	}
	for indexIteratedConnection, iteratedConnection := range chatUserSelected.Connections {
		if iteratedConnection.ID == id {
			chatUserSelected.Connections = append(chatUser.Connections[:indexIteratedConnection], chatUser.Connections[indexIteratedConnection+1])
		}
	}
	return nil
}
