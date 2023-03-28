package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/google/uuid"
)

type InvitationChatUserChatRoomRepositoryInterface interface {
	InsertInvitationChatUserChatRoom(invitationChatUserChatRoom *entity.InvitationChatUserChatRoom) error
	SelectOne(invitationChatUserChatRoom *entity.InvitationChatUserChatRoom, id uuid.UUID) error
	SelectAll(invitationsChatUserChatRoom *[]entity.InvitationChatUserChatRoom) error
	SelectByChatUser(invitationsChatUserChatRoom *[]entity.InvitationChatUserChatRoom, chatUser *entity.ChatUser) error
	SelectByChatRoom(invitationsChatUserChatRoom *[]entity.InvitationChatUserChatRoom, chatRoom *entity.ChatRoom) error
	DeleteInvitationChatUserChatRoom(id uuid.UUID) error
}
