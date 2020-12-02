package concesionaria

import (
	"github.com/jmoiron/sqlx"
	"github.com/NicoLaino/GOSeminario/internal/config"
)

// Message...
type Message struct {
	ID int64
	Text string
}

// ChatService... Public Interface
type ChatService interface {
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}

// Service Struct (not public)
type service struct {
	db *sqlx.DB
	conf *config.Config
}

// New ...
func New (db *sqlx.DB, c *config.Config) (ChatService, error) {
	return service{db, c}, nil
}

func (s service) AddMessage (m Message) error {
	return nil
}

func (s service) FindByID (ID int) *Message {
	return nil
}

func (s service) FindAll () []*Message {
	var list []*Message
	s.db.Select(&list, "SELECT * FROM messages")
	return list
}