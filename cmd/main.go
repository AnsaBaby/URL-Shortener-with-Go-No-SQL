package main

import (
	"log"
	"net/http"
	"url-shortener/internal/shortener"
)

func main() {
	// Setup Cassandra session
	session, err := shortener.NewCassandraSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra: ", err)
	}
	defer session.Close()

	// Initialize repository and service
	repo := shortener.NewURLRepository(session)
	service := shortener.NewURLService(repo)
	handler := shortener.NewHandler(service)

	// Set up routes
	http.HandleFunc("/shorten", handler.ShortenURL)
	http.HandleFunc("/", handler.RedirectURL)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
