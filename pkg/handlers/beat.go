package handlers

import (
	"database/sql"
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
)

type BeatHandler struct {
	DB *sql.DB
}

func (h *BeatHandler) Beat(w http.ResponseWriter, r *http.Request) {

	err := h.DB.Ping()
	if err != nil {
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "database is not running...")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, "server is running...")
}
