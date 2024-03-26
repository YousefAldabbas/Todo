package repository

import (
	"database/sql"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/models"
	"github.com/rs/zerolog/log"
)

type TodoRepo struct {
	DB *sql.DB
}

func (r *TodoRepo) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo

	q := `SELECT * FROM todos`

	rows, err := r.DB.Query(q)
	if err != nil {
		log.Error().Err(err).Msg("Error getting todos")

		return []models.Todo{}, err
	}

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning todos")

			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodoRepo) GetTodoByID(id string) (models.Todo, error) {
	q := `SELECT * FROM todos WHERE id = $1`

	var todo models.Todo

	err := r.DB.QueryRow(q, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepo) InsertTodo(title string) (models.Todo, error) {

	var todo models.Todo
	q := `INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed`
	if err := r.DB.QueryRow(q, title, false).Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepo) UpdateTodoByID(id string, title string, completed bool) (models.Todo, error) {

	q := `UPDATE todos SET title = $1, completed = $2 WHERE id = $3 RETURNING id, title, completed`

	var todo models.Todo
	err := r.DB.QueryRow(q, title, completed, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepo) PatchTodoByID(id string, title *string, completed *bool) (models.Todo, error) {

	q := `UPDATE todos SET title = COALESCE($1, title), completed = COALESCE($2, completed) WHERE id = $3 RETURNING id, title, completed`

	var todo models.Todo
	err := r.DB.QueryRow(q, title, completed, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}


func (r *TodoRepo) DeleteTodoByID(id string) error {
	q := `DELETE FROM todos WHERE id = $1`

	_, err := r.DB.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}