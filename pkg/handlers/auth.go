package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/models"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"github.com/rs/zerolog/log"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	q := `SELECT * FROM users WHERE username = $1`
	var user models.User

	err := h.DB.QueryRow(q, payload.Username).Scan(&user.ID, &user.Username, &user.Sub, &user.Password)

	log.Info().Msgf("User: %v", user)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting user")
		return
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		utils.ResponseWithJSON(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.Sub)
	if err != nil {
		log.Error().Err(err).Msg("Error generating token")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error generating token")
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) TokenValidation(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	claims, err := utils.ValidateToken(token)
	if err != nil {
		utils.ResponseWithJSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, claims)
}
