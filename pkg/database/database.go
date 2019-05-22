package database

import (
	"database/sql"
	"github.com/d7561985/heroku_boilerplate/pkg/config"
	_ "github.com/lib/pq"
	"log"
)

var D *sql.DB

func init() {
	if len(config.V.DatabaseHost) == 0 {
		return
	}
	db, err := sql.Open("postgres", config.V.DatabaseHost)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	D = db
}
