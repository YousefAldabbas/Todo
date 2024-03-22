package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/rs/zerolog/log"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/models"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"github.com/google/uuid"
)

type UserHandler struct {
	DB *sql.DB
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

	q := `INSERT INTO users (username, password, sub) VALUES ($1, $2, $3) RETURNING id`

	if err := h.DB.QueryRow(q, newUser.Username, newUser.Password, newUser.Sub).Scan(&newUser.ID); err != nil {
		log.Error().Err(err).Msg("Error creating user")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	utils.ResponseWithJSON(w, http.StatusCreated, newUser)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	q := `SELECT * FROM users`

	rows, err := h.DB.Query(q)

	if err != nil {
		log.Error().Err(err).Msg("Error getting users")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting users")
		return
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Sub, &user.Username, &user.Password)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning users")
			utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error scanning users")
			return
		}
		users = append(users, user)
	}
	utils.ResponseWithJSON(w, http.StatusOK, users)
}
