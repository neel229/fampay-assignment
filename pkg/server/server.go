package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/neel229/fampay-assignment/pkg/db"
	"github.com/neel229/fampay-assignment/pkg/utils"
)

// NewServer creates a server instance
func NewServer(config *utils.Config, db *db.Store) *Server {
	router := chi.NewRouter()
	return &Server{
		config: config,
		store:  db,
		router: router,
	}
}

// StartServer starts the server
func (s *Server) StartServer() {
	log.Println("starting server on port: ", s.config.Port)
	if err := http.ListenAndServe(":"+s.config.Port, s.router); err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}

func (s *Server) SetupRoutes() {
	s.router.Route("/", func(r chi.Router) {
		r.Post("/", s.YouTubeSearch(context.Background()))
	})
}
