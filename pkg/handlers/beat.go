package handlers

import (
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"net/http"
)




func BeatHandler(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "server is running...")
}
