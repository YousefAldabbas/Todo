package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/models"
	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type TodoHandler struct {
	DB *sql.DB
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	q := `SELECT * FROM todos`

	rows, err := h.DB.Query(q)
	if err != nil {
		log.Error().Err(err).Msg("Error getting todos")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting todos")
		return
	}

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning todos")
			utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error scanning todos")
			return
		}
		todos = append(todos, todo)
	}
	utils.ResponseWithJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "id")

	q := `SELECT * FROM todos WHERE id = $1`

	var todo models.Todo

	err := h.DB.QueryRow(q, todoID).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msg("Todo not found")
			utils.ResponseWithJSON(w, http.StatusNotFound, "Todo not found")
			return
		}
		log.Error().Err(err).Msg("Error getting todo by ID")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error getting todo by ID")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	var body struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg("Error decoding request body")
		utils.ResponseWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	todo.Title = body.Title

	q := `INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed`

	err := h.DB.QueryRow(q, todo.Title, todo.Completed).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		log.Error().Err(err).Msg("Error creating todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error creating todo")
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

	q := `UPDATE todos SET title = COALESCE($1, title), completed = COALESCE($2, completed) WHERE id = $3 RETURNING id, title, completed`

	var todo models.Todo
	err := h.DB.QueryRow(q, body.Title, body.Completed, todoID).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		log.Error().Err(err).Msg("Error updating todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error updating todo")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoID := chi.URLParam(r, "id")

	q := `SELECT exists (SELECT 1 FROM todos WHERE id = $1)`

	var exists bool
	err := h.DB.QueryRow(q, todoID).Scan(&exists)
	if err != nil {
		log.Error().Err(err).Msg("Error checking todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error checking todo")
		return
	}
	if !exists {
		log.Info().Msg("Todo not found")
		utils.ResponseWithJSON(w, http.StatusNotFound, "Todo not found")
		return
	}

	q = `DELETE FROM todos WHERE id = $1`

	_, err = h.DB.Exec(q, todoID)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting todo")
		utils.ResponseWithJSON(w, http.StatusInternalServerError, "Error deleting todo")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, "Todo deleted")

}
