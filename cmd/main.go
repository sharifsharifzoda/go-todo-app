package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"todo_sql_database"
	"todo_sql_database/configs"
	"todo_sql_database/db"
	"todo_sql_database/internal/handler"
	"todo_sql_database/internal/repository"
	"todo_sql_database/internal/service"
	"todo_sql_database/logging"
)

func main() {
	logging.InitLog()
	logger := logging.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Fatal(err)
	}

	if err := InitConfigs(); err != nil {
		logger.Fatalf("error while reading config file. error is %v", err.Error())
	}

	var cfg configs.DatabaseConnConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("Couldn't unmarshal the config into struct. error is %v", err.Error())
	}
	cfg.Password = os.Getenv("DB_PASSWORD")

	conn, err := repository.GetDBConnection(cfg)
	if err != nil {
		logger.Fatalf("error while opening DB. error: %s", err.Error())
	}
	db.Init(conn)

	//---------- Внедрение зависимостей-----------
	newRepository := repository.NewRepository(conn, logger)
	newService := service.NewService(newRepository, logger)
	newHandler := handler.NewHandler(newService.Auth, newService.Todo, logger)
	//--------------------------------------------

	server := new(todo_sql.Server)
	go func() {
		if err := server.Run(os.Getenv("PORT"), newHandler.InitRoutes()); err != nil {
			logger.Fatalf("error while running http.server. Error is %s", err.Error())
		}
	}()
	fmt.Printf("Server is listening to port: %s\n", os.Getenv("PORT"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := conn.Close(); err != nil {
		logger.Fatalf("error while closing DB. Error: %s", err.Error())
	}

	fmt.Println("server is shutting down")
	if err = server.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error while shutting server down. Error: %s", err.Error())
	}
}

func InitConfigs() error {
	viper.AddConfigPath("configs") //адрес директории
	viper.SetConfigName("config")  //имя файла
	viper.SetConfigType("yml")
	return viper.ReadInConfig() //считывает config и сохраняет данные во внутренний объект viper
}
