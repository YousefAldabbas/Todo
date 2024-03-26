package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/repository"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type TodoHandler struct {
	Repo repository.TodoRepo
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {

	todos, err := h.Repo.GetTodos()
	if err != nil {
		log.Error().Err(err).Msg("Error getting todos")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting todos")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "id")

	todo, err := h.Repo.GetTodoByID(todoID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msg("Todo not found")
			utils.ResponseWithJSON(w, http.StatusNotFound, "Todo not found")
			return
		}
		log.Error().Err(err).Msg("Error getting todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting todo")
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	todo, err := h.Repo.InsertTodo(body.Title)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error inserting todo")
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) PatchTodo(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if body.Title == nil && body.Completed == nil {
		log.Error().Msg("No fields to update")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	todoID := chi.URLParam(r, "id")

	todo, err := h.Repo.PatchTodoByID(todoID, body.Title, body.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msg("Todo not found")
			utils.ResponseWithJSON(w, http.StatusNotFound, "Todo not found")
			return
		}
		log.Error().Err(err).Msg("Error updating todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error updating todo")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoID := chi.URLParam(r, "id")

	if err := h.Repo.DeleteTodoByID(todoID); err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msg("Todo not found")
			utils.ResponseWithJSON(w, http.StatusNotFound, "Todo not found")
			return
		}

		log.Error().Err(err).Msg("Error deleting todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error deleting todo")
		return
	}
	utils.ResponseWithJSON(w, http.StatusOK, "Todo deleted")

}
