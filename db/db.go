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

	for i, ddl := range DDls {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("failed to create table #%d due to: %s", i+1, err.Error())
		}
	}
}
