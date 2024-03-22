package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/YousefAldabbas/go-backend-scratch/pkg/handlers"
	"github.com/go-chi/chi/v5"

	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
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

type App struct {
	router *chi.Mux
	DB     *sql.DB
}

func (a *App) New() *App {
	app := &App{
		router: chi.NewRouter(),
		DB:     a.ConnectDB(),
	}

	app.LoadRoutes()
	return app
}

func (a *App) ConnectDB() *sql.DB {
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



func (a *App) Start() {
	http.ListenAndServe(":3000", a.router)
}

func main() {
	app := App{}
	app.New().Start()
}
