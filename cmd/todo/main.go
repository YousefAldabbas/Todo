package main

import (
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/config"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)
type App struct {
	router http.Handler
	DB     *sql.DB

}

func (a *App) New() *App {
	app := &App{
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
