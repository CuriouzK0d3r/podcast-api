package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"podcast-api/database"
	"podcast-api/handlers"
)

func main() {
	// Initialize the database
	database.InitDB("podcasts.db")

	// Create a new router
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/podcasts", handlers.CreatePodcastList).Methods("POST")
	r.HandleFunc("/podcasts", handlers.GetPodcastList).Methods("GET")
	r.HandleFunc("/podcasts/{id}", handlers.DeletePodcast).Methods("DELETE")
	r.HandleFunc("/podcast", handlers.CreatePodcast).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}