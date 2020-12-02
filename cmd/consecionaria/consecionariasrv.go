
package main

import (
	"flag"
	"fmt"
	"os"
	//"time"

	//"github.com/gin-gonic/gin"
	//"github.com/jmoiron/sqlx"
	"github.com/NicoLaino/GOSeminario/internal/config"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	// Catch error if config is not read
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

}