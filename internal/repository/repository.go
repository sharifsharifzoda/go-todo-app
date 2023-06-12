package repository

import (
	"database/sql"
	"todo_sql_database/logging"
	"todo_sql_database/model"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	GetUser(email string) (model.User, error)
	IsEmailUsed(email string) bool
}

type TodoTask interface {
	CreateTask(userId int, task *model.Task) (int, error)
	GetAll(userId int) (model.Tasks, error)
	GetTaskById(userId int, id int) (model.Task, error)
	UpdateTask(userId int, id int, task *model.Task) error
	DeleteTask(userId int, id int) error
}

type Repository struct {
	Authorization
	TodoTask
	Logger *logging.Logger
}

func NewRepository(db *sql.DB, log *logging.Logger) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoTask:      NewTodoTaskPostgres(db, log),
		Logger:        log,
	}
}
