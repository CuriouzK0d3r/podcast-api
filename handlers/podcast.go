package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"podcast-api/models"
	"github.com/gorilla/mux"
)

// CreatePodcastList handles the POST request to store a list of podcasts
func CreatePodcastList(w http.ResponseWriter, r *http.Request) {
	var podcasts []models.Podcast
	if err := json.NewDecoder(r.Body).Decode(&podcasts); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, podcast := range podcasts {
		if err := podcast.Save(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

// GetPodcastList handles the GET request to retrieve the list of podcasts
func GetPodcastList(w http.ResponseWriter, r *http.Request) {
	podcasts, err := models.FetchAllPodcasts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(podcasts)
}

// DeletePodcast handles the DELETE request to remove a specific podcast
func DeletePodcast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "Missing podcast ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid podcast ID", http.StatusBadRequest)
		return
	}

	if err := models.DeletePodcast(id); err != nil {
		if err.Error() == fmt.Sprintf("podcast with id %d not found", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreatePodcast handles the POST request to store a single podcast
func CreatePodcast(w http.ResponseWriter, r *http.Request) {
	var podcast models.Podcast
	if err := json.NewDecoder(r.Body).Decode(&podcast); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Basic validation
	if podcast.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if podcast.Link == "" {
		http.Error(w, "Link is required", http.StatusBadRequest)
		return
	}

	if err := podcast.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(podcast)
}