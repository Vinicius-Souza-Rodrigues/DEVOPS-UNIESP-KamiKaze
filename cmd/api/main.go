package main

import (
	"log"
	"net/http"

	"github.com/uniesp/desafio-devops/internal/database"
	"github.com/uniesp/desafio-devops/internal/handler"
	"github.com/uniesp/desafio-devops/internal/repository"
	"github.com/uniesp/desafio-devops/internal/service"
)

func main() {

	db := database.Connect()

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/health", handler.Health)

	http.HandleFunc(
		"/users",
		userHandler.Create,
	)

	log.Println("server rodando na porta 8080")

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
