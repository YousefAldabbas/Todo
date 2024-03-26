package main

import (
	"github.com/YousefAldabbas/go-backend-scratch/pkg/handlers"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func (a *App) LoadTodoRoutes(r chi.Router) {

	h := handlers.TodoHandler{
		Repo: repository.TodoRepo{
			DB: a.DB,
		},
	}
	r.Get("/", h.GetTodos)
	r.Get("/{id}", h.GetTodoByID)

	r.Post("/", h.CreateTodo)
	r.Patch("/{id}", h.PatchTodo)
	r.Delete("/{id}", h.DeleteTodo)
}

func (a *App) LoadUserRoutes(r chi.Router) {
	h := &handlers.UserHandler{
		Repo: repository.UserRepo{
			DB: a.DB,
		},
	}

	r.Get("/", h.GetAllUsers)
	r.Post("/", h.CreateUser)
	r.Get("/{sub}", h.GetUserBySub)
}

func (a *App) LoadAuthRoutes(r chi.Router) {
	h := handlers.AuthHandler{DB: a.DB}
	r.Post("/login", h.Login)
	// for testing purposes
	r.Get("/token-validation", h.TokenValidation)
}

func (a *App) LoadBeatRoutes(r chi.Router) {

	h := handlers.BeatHandler{DB: a.DB}
	r.Get("/", h.Beat)
}

func (a *App) LoadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// private routes
	r.Group(func(r chi.Router) {
		tokenAuth := jwtauth.New("HS256", []byte("secret-key"), nil)
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/beat", a.LoadBeatRoutes)
	})

	r.Route("/todos", a.LoadTodoRoutes)
	r.Route("/users", a.LoadUserRoutes)
	r.Route("/auth", a.LoadAuthRoutes)

	a.router = r
}
