package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"todo_sql_database/configs"
)

func GetDBConnection(cfg configs.DatabaseConnConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
	//log.Print(dsn)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("Connection success host:%s port:%s", cfg.Host, cfg.Port)

	return conn, nil
}
