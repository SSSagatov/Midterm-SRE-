package server

import (
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	address        string
	port           uint16
	fileServers    []*FileServer
	registerRoutes func(mux *http.ServeMux, handler interface{})
}

func New(address string, port uint16, fileServers []*FileServer, registerRoutes func(mux *http.ServeMux, handler interface{})) *Server {
	return &Server{
		address:        address,
		port:           port,
		fileServers:    fileServers,
		registerRoutes: registerRoutes,
	}
}

func (s *Server) Run(handler interface{}) {
	mux := http.NewServeMux()
	for _, fileServer := range s.fileServers {
		fs := http.FileServer(http.Dir(fileServer.Path))
		mux.Handle("GET "+fileServer.Url, http.StripPrefix(fileServer.Url, fs))
	}

	s.registerRoutes(mux, handler)

	addr := s.address + ":" + strconv.Itoa(int(s.port))
	log.Printf("Starting server at %s...", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
