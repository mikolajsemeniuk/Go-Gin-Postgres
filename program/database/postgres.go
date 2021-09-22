package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgres_username = "postgres_username"
	postgres_password = "postgres_password"
	postgres_host     = "postgres_host"
	postgres_port     = "postgres_port"
	postgres_database = "postgres_database"
)

var (
	Client *sql.DB

	username = os.Getenv(postgres_username)
	password = os.Getenv(postgres_password)
	host     = os.Getenv(postgres_host)
	port     = os.Getenv(postgres_port)
	database = os.Getenv(postgres_database)
)

func init() {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	// let this error like this otherwise Client would be declared once again and throw nil dereference exception
	var error error
	Client, error = sql.Open("postgres", dataSourceName)
	if error != nil {
		panic(error)
	}

	if error = Client.Ping(); error != nil {
		panic(error)
	}
}
