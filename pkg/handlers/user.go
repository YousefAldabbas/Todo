package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/models"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/repository"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"github.com/google/uuid"
)

type UserHandler struct {
	Repo repository.UserRepo
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing the password")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Error while hashing the password")
		return
	}

	newUser.Username = payload.Username
	newUser.Password = hashedPassword
	newUser.Sub = uuid.New().String()

	newUser, err = h.Repo.InsertUser(newUser)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting user")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error inserting user")
		return
	}

	utils.ResponseWithJSON(w, http.StatusCreated, newUser)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAllUsers()

	if err != nil {
		log.Error().Err(err).Msg("Error getting users")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting users")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, users)
}



func (h *UserHandler) GetUserBySub(w http.ResponseWriter, r *http.Request) {
	sub := chi.URLParam(r, "sub")

	user, err := h.Repo.GetUserBySub(sub)
	if err != nil {
		if err == sql.ErrNoRows{
			log.Info().Msg("User not found")
			utils.ResponseWithJSON(w, http.StatusNotFound, "User not found")
		}
		log.Error().Err(err).Msg("Error getting user")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting user")
	}
	utils.ResponseWithJSON(w, http.StatusOK, user)
}