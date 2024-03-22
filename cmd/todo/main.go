package main

import (
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/config"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"github.com/go-chi/chi/v5"

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
	utils.InitLogger()
	return app
}

func (a *App) Start() {
	http.ListenAndServe(":3000", a.router)
}

func main() {
	app := App{}
	app.New().Start()
}
