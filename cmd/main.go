package main

import (
	"log"

	"github.com/neel229/fampay-assignment/pkg/db"
	"github.com/neel229/fampay-assignment/pkg/server"
	"github.com/neel229/fampay-assignment/pkg/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}
	store, err := db.NewStore(config.DBURL)
	if err != nil {
		log.Fatalf("error instantiating db conn: %v\n", err)
	}
	svr := server.NewServer(config, store)
	svr.SetupRoutes()
	svr.StartServer()
}
