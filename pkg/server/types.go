package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/neel229/fampay-assignment/pkg/db"
	"github.com/neel229/fampay-assignment/pkg/utils"
)

// Server holds instances of config,
// db conn and a router
type Server struct {
	config *utils.Config
	store  *db.Store
	router *chi.Mux
}
