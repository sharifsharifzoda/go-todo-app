package repository

import (
	"database/sql"
	"todo_sql_database/db"
	"todo_sql_database/model"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *model.User) (int, error) {
	var id int
	row := r.db.QueryRow(db.CreateUser, user.Name, user.Email, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow(db.GetUser, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsActive,
	)
	return user, err
}

func (r *AuthPostgres) IsEmailUsed(email string) bool {
	var user model.User
	row := r.db.QueryRow(db.SelectEmail, email)
	err := row.Scan(&user.Email)
	if err != nil {
		return false
	}
	return true
}
