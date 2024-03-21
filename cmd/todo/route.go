package main

import (
	"github.com/YousefAldabbas/go-backend-scratch/pkg/handlers"

	"net/http"

	"github.com/go-chi/chi/v5"
)

//  Routes function
//  This function is used to define the routes for the books resource

func BooksRoutes() http.Handler {
	r := chi.NewRouter()
	h := &handlers.BooksHandler{}

	r.Get("/", h.GetBooks)
	r.Get("/{id}", h.GetBook)

	return r
}
