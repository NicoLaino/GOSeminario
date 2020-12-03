package concesionaria

import (
	"github.com/jmoiron/sqlx"
	"github.com/NicoLaino/GOSeminario/internal/config"
	"fmt"
)

// Message...
type Message struct {
	ID int64
	Text string
}

type queryResult struct {
	TextResult string
}

// Service Public Interface
type Service interface {
	AddMessage(Message) (*queryResult, error) 
	FindByID(int) (*Message, error)
	FindAll() ([]*Message, error)
	DeleteByID(int) (*queryResult, error)
}

// Service Struct (not public)
type service struct {
	db *sqlx.DB
	conf *config.Config
}

// New ...
func New (db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddMessage (message Message) (*queryResult, error) {
	sqlStatement := "INSERT INTO messages (text) VALUES (:text);"

	result, err := s.db.NamedExec(sqlStatement, &message)

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	sqlResult := &queryResult{
		TextResult: "Inserted Row.",
	}

	return sqlResult, nil
}

func (s service) FindAll () ([]*Message, error) {
	var list []*Message
	
	if err := s.db.Select(&list, "SELECT * FROM messages");err != nil {
		return nil, err
	}

	return list, nil
}


func (s service) FindByID(ID int) (*Message, error) {
	var message Message
	sqlStatement := "SELECT * FROM messages WHERE id=?;"

	err := s.db.QueryRowx(sqlStatement, ID).StructScan(&message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (s service) DeleteByID(ID int) (*queryResult, error) {
	sqlStatement := "DELETE FROM messages WHERE id=?;"

	result, err := s.db.Exec(sqlStatement, ID)

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	sqlResult := &queryResult{
		TextResult: "Deleted Row.",
	}

	return sqlResult, nil
}