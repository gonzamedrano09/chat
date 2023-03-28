package entity

type ChatUser struct {
	User        *User
	Connections []Connection
}
