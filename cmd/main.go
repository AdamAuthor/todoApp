package main

import (
	"github.com/spf13/viper"
	"log"
	"todoApp"
	"todoApp/pkg/handler"
	"todoApp/pkg/repository"
	"todoApp/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with reading configs: %s", err.Error())
	}
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todoApp.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error with runnig sеrver: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
