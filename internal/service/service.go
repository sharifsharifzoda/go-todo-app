package service

import (
	"todo_sql_database/internal/repository"
	"todo_sql_database/logging"
	"todo_sql_database/model"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	ValidateUser(user model.User) error
	IsEmailUsed(email string) bool
	CreateUser(user *model.User) (int, error)
	CheckUser(user model.User) (model.User, error)
	GenerateToken(user model.User) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoTask interface {
	CreateTask(userId int, task *model.Task) (int, error)
	GetAll(userId int) (model.Tasks, error)
	GetTaskById(userId int, id int) (model.Task, error)
	ValidateTask(task *model.Task) error
	UpdateTask(userId int, id int, task *model.Task) error
	DeleteTask(userId int, id int) error
}

type Service struct {
	Auth   Authorization
	Todo   TodoTask
	Logger *logging.Logger
}

func NewService(repos *repository.Repository, log *logging.Logger) *Service {
	return &Service{
		Auth:   NewAuthService(repos.Authorization, log),
		Todo:   NewTodoTaskService(repos.TodoTask, log),
		Logger: log,
	}
}
