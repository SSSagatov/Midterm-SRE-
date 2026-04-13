package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"campus-connect/backend/internal/db"
	"campus-connect/backend/internal/handlers"
	"campus-connect/backend/internal/middleware"
	"campus-connect/backend/internal/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := getenv("APP_PORT", "8080")
	databaseURL := getenv("DATABASE_URL", "postgres://campus_user:campus_pass@localhost:5432/campus_connect?sslmode=disable")

	ctx := context.Background()
	pool, err := db.NewPool(ctx, databaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()

	repo := db.NewPostRepository(pool)
	appService := service.NewPostService(repo)
	h := handlers.NewHandler(appService)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.Health)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/api/posts", h.Posts)
	mux.HandleFunc("/api/posts/", h.PostAction)
	mux.HandleFunc("/debug/fail", h.ForceFail)
	mux.HandleFunc("/debug/slow", h.ForceSlow)

	wrapped := middleware.WithCORS(middleware.WithMetrics(mux))

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           wrapped,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("backend listening on :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
