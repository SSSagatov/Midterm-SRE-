package ssr

import (
	"log"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler interface{}) {
	h, ok := handler.(*Handler)
	if !ok {
        log.Fatalf("Expected handler of type *Handler, got %T", handler)
        return
    }
	mux.HandleFunc("GET /", h.Index)
	mux.HandleFunc("GET /artist", h.Artist)
	mux.HandleFunc("/", h.NotFound)
}