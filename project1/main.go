package main

import (
	"go-user-app/repository"
	"go-user-app/service"
	"project1/handler"
)

func main() {
	// Choose implementation
	repo := repository.PostgresRepo{}
	// repo := repository.MockRepo{}

	userService := service.UserService{Repo: repo}
	userHandler := handler.UserHandler{Service: userService}

	userHandler.GetUser(1)
}
