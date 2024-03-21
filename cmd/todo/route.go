package main

import (
	"github.com/YousefAldabbas/scratch-chi/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"net/http"
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
