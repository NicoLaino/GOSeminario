package concesionaria

import (
	"github.com/jmoiron/sqlx"
	"github.com/NicoLaino/GOSeminario/internal/config"
	"fmt"
)

// Car...
type Car struct {
	ID int
	Text string
}

type queryResult struct {
	TextResult string
}

// Service Public Interface
type Service interface {
	AddCar(Car) (*queryResult, error) 
	FindByID(int) (*Car, error)
	FindAll() ([]*Car, error)
	DeleteByID(int) (*queryResult, error)
	UpdateByID(int, Car) (*queryResult, error)
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

func (s service) AddCar (car Car) (*queryResult, error) {
	sqlStatement := "INSERT INTO cars (text) VALUES (:text);"

	result, err := s.db.NamedExec(sqlStatement, &car)

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	sqlResult := &queryResult{
		TextResult: "Inserted Row.",
	}

	return sqlResult, nil
}

func (s service) FindAll () ([]*Car, error) {
	var list []*Car
	
	if err := s.db.Select(&list, "SELECT * FROM cars");err != nil {
		return nil, err
	}

	return list, nil
}


func (s service) FindByID(ID int) (*Car, error) {
	var car Car
	sqlStatement := "SELECT * FROM cars WHERE id=?;"

	err := s.db.QueryRowx(sqlStatement, ID).StructScan(&car)

	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (s service) DeleteByID(ID int) (*queryResult, error) {
	sqlStatement := "DELETE FROM cars WHERE id=?;"

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

func (s service) UpdateByID(ID int, car Car) (*queryResult, error) {

	car.ID = ID

	sqlStatement := `UPDATE cars SET text=:text WHERE id=:id;`

	result, err := s.db.NamedExec(sqlStatement, &car)

	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	sqlResult := &queryResult{
		TextResult: "Row Updated.",
	}

	return sqlResult, nil
}