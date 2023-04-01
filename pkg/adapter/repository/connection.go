package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
	"github.com/google/uuid"
)

type ConnectionRepository struct {
	Connections []entity.Connection
}

func NewConnectionRepository() repository.ConnectionRepositoryInterface {
	return &ConnectionRepository{}
}

func (cr *ConnectionRepository) InsertConnection(connection *entity.Connection) error {
	cr.Connections = append(cr.Connections, *connection)
	return nil
}

func (cr *ConnectionRepository) SelectOne(connection *entity.Connection, id uuid.UUID) error {
	for _, iteratedConnection := range cr.Connections {
		if iteratedConnection.ID == id {
			connection = &iteratedConnection
		}
	}
	return nil
}

func (cr *ConnectionRepository) SelectAll(connections *[]entity.Connection) error {
	connections = &cr.Connections
	return nil
}

func (cr *ConnectionRepository) DeleteConnection(id uuid.UUID) error {
	for indexIteratedConnection, iteratedConnection := range cr.Connections {
		if iteratedConnection.ID == id {
			cr.Connections = append(cr.Connections[:indexIteratedConnection], cr.Connections[indexIteratedConnection+1])
		}
	}
	return nil
}
