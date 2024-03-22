package handlers

import (
	"database/sql"
	"net/http"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/utils"
)

type TodoHandler struct{
	DB     *sql.DB
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "GetTodos")
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "GetTodoByID")
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "CreateTodo")
}

func (h *TodoHandler) UpdateTodoByID(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "UpdateTodoByID")
}

func (h *TodoHandler) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "DeleteTodoByID")
}
