package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) GetVideos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitParam := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			log.Println(err)
			http.Error(w, "error decoding limit value", http.StatusBadRequest)
			return
		}
		offsetParam := r.URL.Query().Get("offset")
		offset, err := strconv.Atoi(offsetParam)
		if err != nil {
			log.Println(err)
			http.Error(w, "error decoding offset value", http.StatusBadRequest)
			return
		}
		videos, err := s.store.GetVideos(limit, offset)
		if err != nil {
			log.Println(err)
			http.Error(w, "error fetching videos, try again later...", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(videos)
	}
}

func (s *Server) SearchWithTitle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
		video, err := s.store.SearchWithTitle(input.Title)
		if err != nil {
			http.Error(w, "video not found", http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(video)
	}
}

func (s *Server) SearchWithDescription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Description string `json:"description"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
		video, err := s.store.SearchWithDescription(input.Description)
		if err != nil {
			http.Error(w, "video not found", http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(video)
	}
}
