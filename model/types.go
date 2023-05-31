package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	IsActive    bool      `json:"is_active"`
	Deadline    string    `json:"deadline,omitempty"`
	UserId      int       `json:"-"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type Tasks []Task
