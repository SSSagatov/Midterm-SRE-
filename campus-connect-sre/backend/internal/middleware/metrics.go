package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path", "status"},
	)

	postCreateTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "post_create_total",
			Help: "Total post creation attempts",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(requestTotal, requestDuration, postCreateTotal)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func WithMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec, r)

		path := normalizePath(r.URL.Path)
		status := strconv.Itoa(rec.status)
		requestTotal.WithLabelValues(r.Method, path, status).Inc()
		requestDuration.WithLabelValues(r.Method, path, status).Observe(time.Since(start).Seconds())
	})
}

func ObservePostCreate(status string) {
	postCreateTotal.WithLabelValues(status).Inc()
}

func normalizePath(path string) string {
	switch {
	case path == "/api/posts":
		return "/api/posts"
	case len(path) >= 11 && path[:11] == "/api/posts/":
		return "/api/posts/:id/action"
	default:
		return path
	}
}
