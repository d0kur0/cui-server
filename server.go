package main

import (
	"log"

	"github.com/d0kur0/cui-server/database"

	"github.com/d0kur0/cui-server/config"
	"github.com/d0kur0/cui-server/server"
)

func main() {
	appConfig := config.New()
	if err := appConfig.Init(); err != nil {
		log.Fatalf("Startup error on init appConfig; %s", err)
	}

	if err := database.Init(); err != nil {
		log.Fatalf("Startup error on init database; %s", err)
	}

	appServer := server.New()
	if err := appServer.Init(appConfig); err != nil {
		log.Fatalf("Startup error on init appServer; %s", err)
	}
}
