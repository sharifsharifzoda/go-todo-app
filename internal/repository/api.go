package repository

import (
	"database/sql"
	"log"
	"todo_sql_database/db"
	"todo_sql_database/model"
)

type TodoTaskPostgres struct {
	db *sql.DB
}

func NewTodoTaskPostgres(db *sql.DB) *TodoTaskPostgres {
	return &TodoTaskPostgres{db: db}
}

func (r *TodoTaskPostgres) CreateTask(userId int, task *model.Task) (int, error) {
	var id int
	row := r.db.QueryRow(db.CreateTask, task.Name, task.Description, task.Deadline, userId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TodoTaskPostgres) GetAll(userId int) (model.Tasks, error) {
	var tasks model.Tasks

	rows, err := r.db.Query(db.GetAllTasks, userId)
	if err != nil {
		log.Printf("failed to get row. error is %v", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.Done, &t.IsActive, &t.Deadline, &t.Username)
		if err != nil {
			log.Println("failed to scan due to:", err.Error())
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TodoTaskPostgres) GetTaskById(userId int, id int) (model.Task, error) {
	var t model.Task

	row := r.db.QueryRow(db.GetTaskById, id, userId)
	err := row.Scan(&t.Id, &t.Name, &t.Description, &t.Done, &t.IsActive, &t.Deadline, &t.Username)
	if err != nil {
		return model.Task{}, err
	}

	return t, nil
}

func (r *TodoTaskPostgres) UpdateTask(userId int, id int, task *model.Task) error {
	_, err := r.db.Exec(db.UpdateTask, task.Name, task.Description, task.Done, task.IsActive, task.Deadline, id, userId)
	return err
}

func (r *TodoTaskPostgres) DeleteTask(userId int, id int) error {
	_, err := r.db.Exec(db.DeleteTask, id, userId)
	return err
}
