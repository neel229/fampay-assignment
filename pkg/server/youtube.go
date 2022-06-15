package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/neel229/fampay-assignment/pkg/db"
	yt "github.com/neel229/fampay-assignment/pkg/youtube"
)

// YouTubeSearch takes the keyword input from request,
// searches for latest videos and stores in db
func (s *Server) YouTubeSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sq yt.SearchQuery
		if err := json.NewDecoder(r.Body).Decode(&sq); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
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
