package ssr

import (
	"errors"
	"groupie-tracker/internal/application"
	"groupie-tracker/internal/domain"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.NotFound(w, r)
		return
	}
	artists, err := h.service.GetArtists()
	if err != nil {
		log.Println("Error fetching artists:", err)
		h.IntrenalServerError(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("internal/infrastructure/assets/html/index.html"))
	tmpl.Execute(w, artists)
}

func (h *Handler) Artist(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.BadRequest(w, r)
		return
	}
	artist, err := h.service.GetArtist(id)
	if errors.Is(err, domain.ErrIncorrectId) {
		h.BadRequest(w, r)
		return
	}
	if err != nil {
		log.Println("Error fetching artist:", err)
		h.IntrenalServerError(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("internal/infrastructure/assets/html/artist.html"))
	tmpl.Execute(w, artist)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("internal/infrastructure/assets/html/404.html"))
	tmpl.Execute(w, "")
}

func (h *Handler) BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	tmpl := template.Must(template.ParseFiles("internal/infrastructure/assets/html/400.html"))
	tmpl.Execute(w, "")
}

func (h *Handler) IntrenalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	tmpl := template.Must(template.ParseFiles("internal/infrastructure/assets/html/500.html"))
	tmpl.Execute(w, "")
}
