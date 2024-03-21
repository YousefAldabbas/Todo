package handlers

import (
	"github.com/YousefAldabbas/scratch-chi/pkg/utils"
)


type BooksHandler struct{}

func (h *BooksHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "get books...")
}

func (h *BooksHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "get book...")
}
