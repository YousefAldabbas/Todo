package config

import (
	"database/sql"
	"fmt"
	"log"
)

type DBConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func (c DBConfig) ConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database)
}
func ConnectDB() *sql.DB {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "todo",
		User:     "postgres",
		Password: "postgres",
	}

	db, err := sql.Open("pgx", dbConfig.ConnString())
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	return db
}
