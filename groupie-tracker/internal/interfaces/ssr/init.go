package ssr

import (
	"groupie-tracker/internal/application"
	"groupie-tracker/internal/infrastructure/server"
)

func Init(service *application.Service) {
	fs := server.NewFileServer("./internal/infrastructure/assets", "/static/")
	fss := []*server.FileServer{fs}
	server := server.New("localhost", 8080, fss, RegisterRoutes)
	handler := NewHandler(service)
	server.Run(handler)
}
