package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/google/uuid"
)

type ConnectionRepositoryInterface interface {
	InsertConnection(connection *entity.Connection) error
	SelectOne(connection *entity.Connection, id uuid.UUID) error
	SelectAll(connections *[]entity.Connection) error
	DeleteConnection(id uuid.UUID) error
}
