package db

import (
	"database/sql"
	"log"
)

func Init(db *sql.DB) {
	DDls := []string{
		CreateUsersTable,
		CreateTasksTable,
	}

	for _, ddl := range DDls {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("failed to create tables due to: %s", err.Error())
		}
	}
}
