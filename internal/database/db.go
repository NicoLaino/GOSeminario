package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/NicoLaino/GOSeminario/internal/config"
	_ "github.com/mattn/go-sqlite3" // SQLlite3 driver support
)

// Permite instanciar distintos tipos de DB, por eso se le define que tipo
// New Database ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type{
		case "sqlite3" :
			db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
			if err != nil {
				return nil, err
			}

			db.Ping()
			if err != nil {
				return nil, err
			}

			return db, nil
		default: 
			return nil, errors.New("Invalid DB Type")
	}
}

