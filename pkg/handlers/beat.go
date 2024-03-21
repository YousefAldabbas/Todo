package handlers

import (
	"net/http"
	"github.com/yourusername/chi-todo/pkg/utils"
)



func BeatHandler(w http.ResponseWriter, r *http.Request) {
	ResponseWithJSON(w, http.StatusOK, "server is running...")
}