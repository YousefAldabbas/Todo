package main

import (
	"github.com/YousefAldabbas/go-backend-scratch/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) LoadTodoRoutes(router chi.Router) {

	h := handlers.TodoHandler{DB: a.DB}
	router.Get("/", h.GetTodos)
	router.Get("/{id}", h.GetTodoByID)
}

func (a *App) LoadBeatRoutes(router chi.Router) {
	h := handlers.BeatHandler{DB: a.DB}
	router.Get("/", h.Beat)
}

func (a *App) LoadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/beat", a.LoadBeatRoutes)
	r.Route("/todos", a.LoadTodoRoutes)
	a.router = r
}
