package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// background search for football keyword
	// with an interval of 10 seconds
	tickRate := time.Second * time.Duration(config.TickRate)
	go bgSearch(tickRate, "football", svr)
	svr.StartServer()
}

func bgSearch(tickRate time.Duration, keyword string, svr *server.Server) {
	svr.YouTubeServerSearch(keyword)
	ticker := time.NewTicker(tickRate).C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-ticker:
			svr.YouTubeServerSearch(keyword)
		case <-c:
			os.Exit(1)
		}
	}
}
