package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/neel229/fampay-assignment/pkg/db"
	yt "github.com/neel229/fampay-assignment/pkg/youtube"
)

// YouTubeSearch takes the keyword input from request,
// searches for latest videos and stores in db
func (s *Server) YouTubeSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sq yt.SearchQuery
		json.NewDecoder(r.Body).Decode(&sq)
		videos, err := yt.Search(s.config.GoogleAPIKey, sq.Keyword)
		if err != nil {
			http.Error(w, yt.ErrSearchingVideos.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(videos)
		ra, err := s.store.InsertVideos(videos)
		if err != nil {
			http.Error(w, db.ErrInsertingVideo.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(ra)
	}
}

// YouTubeServerSearch searches for latest videos and stores in db
func (s *Server) YouTubeServerSearch(keyword string) error {
	videos, err := yt.Search(s.config.GoogleAPIKey, keyword)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.store.InsertVideos(videos)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

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
