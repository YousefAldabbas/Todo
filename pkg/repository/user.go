package repository

import (
	"database/sql"

	"github.com/YousefAldabbas/go-backend-scratch/pkg/models"
	"github.com/rs/zerolog/log"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) InsertUser(user models.User) (models.User, error) {
	q := `INSERT INTO users (username, password, sub) VALUES ($1, $2, $3) RETURNING id`
	if err := r.DB.QueryRow(q, user.Username, user.Password, user.Sub).Scan(&user.ID); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetAllUsers() ([]models.User, error) {

	var users []models.User
	q := `SELECT * FROM users`

	rows, err := r.DB.Query(q)

	if err != nil {
		log.Error().Err(err).Msg("Error getting users")
		return nil, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Sub, &user.Username, &user.Password)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning users")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}


func (r *UserRepo) GetUserBySub(sub string) (models.User, error) {
	var user models.User
	q := `SELECT * FROM users WHERE sub = $1`

	if err := r.DB.QueryRow(q, sub).Scan(&user.ID, &user.Sub, &user.Username, &user.Password); err != nil {
		return models.User{}, err
	}
	return user, nil
}