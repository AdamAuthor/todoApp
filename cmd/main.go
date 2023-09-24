package main

import (
	"log"
	"todoApp"
	"todoApp/pkg/handler"
	"todoApp/pkg/repository"
	"todoApp/pkg/service"
)

func main() {

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todoApp.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error with runnig srver: %s", err.Error())
	}

}
