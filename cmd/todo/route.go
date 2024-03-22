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

	router.Post("/", h.CreateTodo)
	router.Patch("/{id}", h.PatchTodo)
	router.Delete("/{id}", h.DeleteTodo)
}

func (a *App) LoadUserRoutes(router chi.Router) {
	h := &handlers.UserHandler{
		DB: a.DB,
	}

	router.Get("/", h.GetAllUsers)
	router.Post("/", h.CreateUser)
}

func (a *App) LoadAuthRoutes(router chi.Router) {
	h := handlers.AuthHandler{DB: a.DB}
	router.Post("/login", h.Login)
	// for testing purposes
	router.Get("/token-validation", h.TokenValidation)
}

func (a *App) LoadBeatRoutes(router chi.Router) {
	h := handlers.BeatHandler{DB: a.DB}
	router.Get("/", h.Beat)
}

func (a *App) LoadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/beat", a.LoadBeatRoutes)
	r.Route("/todos", a.LoadTodoRoutes)
	r.Route("/users", a.LoadUserRoutes)
	r.Route("/auth", a.LoadAuthRoutes)
	
	a.router = r
}
