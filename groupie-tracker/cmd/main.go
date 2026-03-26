package main

import (
	"groupie-tracker/internal/application"
	"groupie-tracker/internal/infrastructure/persistence/api"
	"groupie-tracker/internal/interfaces/ssr"
)

func main() {
	repo := api.New("https://groupietrackers.herokuapp.com")
	service := application.NewService(repo)
	ssr.Init(service)
}
