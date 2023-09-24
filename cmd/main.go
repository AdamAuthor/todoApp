package main

import (
	"log"
	"todoApp"
	"todoApp/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todoApp.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error with runnig srver: %s", err.Error())
	}

}
