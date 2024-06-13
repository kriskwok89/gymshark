package main

import (
	"log"
	"net/http"

	"gymshark-challenge/internal/config"
	"gymshark-challenge/internal/handlers"
	"gymshark-challenge/internal/services"
)

func main() {
	cfg := config.LoadConfig("config/config.properties")
	calculator := services.NewPackCalculator(cfg.PackSizes)
	handler := handlers.NewPacksHandler(calculator)

	// Serve static files from Vue build directory
	fs := http.FileServer(http.Dir("./vue-app/dist"))
	http.Handle("/", logRequest(fs))

	// API endpoint
	http.Handle("/api/packs", logRequest(handler))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// logRequest logs details of each incoming HTTP request
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}
