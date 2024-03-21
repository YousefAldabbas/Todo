package main

import (
	"github.com/YousefAldabbas/go-backend-scratch/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the application...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/beat", handlers.BeatHandler)
	r.Mount("/books", BooksRoutes())

	http.ListenAndServe(":3000", r)

}
