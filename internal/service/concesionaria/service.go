package concesionaria

import "github.com/NicoLaino/GOSeminario/internal/config"

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
	conf *config.Config
}

// New ...
func New (c *config.Config) (ChatService, error) {
	return service{c}, nil
}

func (s service) AddMessage (m Message) error {
	return nil
}

func (s service) FindByID (ID int) *Message {
	return nil
}

func (s service) FindAll () []*Message {
	var list []*Message
	list = append(list, &Message{0, "Hellow World"})
	return list
}