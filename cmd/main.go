package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"todoApp"
	"todoApp/pkg/handler"
	"todoApp/pkg/repository"
	"todoApp/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with reading configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		fmt.Errorf("error with loading env files: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todoApp.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error with runnig sеrver: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
