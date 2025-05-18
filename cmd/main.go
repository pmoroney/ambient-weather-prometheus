package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pmoroney/ambient-weather-prometheus/internal/handlers"
	"github.com/pmoroney/ambient-weather-prometheus/internal/metrics"
)

// LoggingMiddleware logs details about each incoming HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Received request: Method=%s, URL=%s, RemoteAddr=%s", r.Method, r.URL.String(), r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Request processed: Method=%s, URL=%s, Duration=%s", r.Method, r.URL.String(), time.Since(start))
	})
}

func main() {
	// Initialize Prometheus metrics
	metrics.RegisterMetrics()

	// Create a new router
	r := mux.NewRouter()

	// Add the logging middleware
	r.Use(LoggingMiddleware)

	// Define the webhook endpoint
	r.HandleFunc("/webhook", handlers.WebhookHandler).Methods("GET")

	// Define the metrics endpoint
	r.HandleFunc("/metrics", metrics.ExposeMetrics).Methods("GET")

	// Start the HTTP server
	log.Println("Starting server on :9876")
	if err := http.ListenAndServe(":9876", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
