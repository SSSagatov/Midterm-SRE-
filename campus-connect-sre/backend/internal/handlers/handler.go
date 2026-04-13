package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"campus-connect/backend/internal/middleware"
	"campus-connect/backend/internal/models"
	"campus-connect/backend/internal/service"
)

type Handler struct {
	service *service.PostService
}

func NewHandler(service *service.PostService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"service": "campus-connect-backend",
	})
}

func (h *Handler) Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listPosts(w, r)
	case http.MethodPost:
		h.createPost(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *Handler) PostAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 4 || parts[0] != "api" || parts[1] != "posts" || parts[3] != "like" {
		writeError(w, http.StatusNotFound, "route not found")
		return
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	if err := h.service.LikePost(r.Context(), id); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "post liked"})
}

func (h *Handler) ForceFail(w http.ResponseWriter, _ *http.Request) {
	writeError(w, http.StatusInternalServerError, "intentional failure for alert validation")
}

func (h *Handler) ForceSlow(w http.ResponseWriter, r *http.Request) {
	seconds, _ := strconv.Atoi(r.URL.Query().Get("seconds"))
	if seconds <= 0 {
		seconds = 6
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "slow endpoint executed",
	})
}

func (h *Handler) listPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.ListPosts(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, posts)
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var req models.Post
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middleware.ObservePostCreate("bad_request")
		writeError(w, http.StatusBadRequest, "invalid json payload")
		return
	}

	created, err := h.service.CreatePost(ctx, req)
	if err != nil {
		middleware.ObservePostCreate("failed")
		status := http.StatusInternalServerError
		if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		} else if strings.Contains(err.Error(), "required") {
			status = http.StatusBadRequest
		}
		writeError(w, status, err.Error())
		return
	}

	middleware.ObservePostCreate("success")
	writeJSON(w, http.StatusCreated, created)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
