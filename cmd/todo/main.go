package main

import (
	"github.com/YousefAldabbas/go-backend-scratch/pkg/config"
	"github.com/go-chi/chi/v5"
	"net/http"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	router *chi.Mux
	DB     *sql.DB
}

func (a *App) New() *App {
	app := &App{
		router: chi.NewRouter(),
		DB:     config.ConnectDB(),
	}

	app.LoadRoutes()
	return app
}

func (a *App) Start() {
	http.ListenAndServe(":3000", a.router)
}

func main() {
	app := App{}
	app.New().Start()
}
