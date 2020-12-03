package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/NicoLaino/GOSeminario/internal/config"
	"github.com/NicoLaino/GOSeminario/internal/database"
	"github.com/NicoLaino/GOSeminario/internal/service/concesionaria"
)

func main() {
	// Leer configuración
	cfg := readConfig()

	// Crear Conexión a DB
	db, err := database.NewDatabase(cfg)
	defer db.Close()
	
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Instanciar servicio e injectar db y configuración
	service, _ := concesionaria.New(db, cfg)
	httpService := concesionaria.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()

}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	// Catch error if config is not read
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primary key autoincrement,
		text varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}