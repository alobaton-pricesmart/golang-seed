package main

import (
	"golang-seed/apps/core/pkg/config"
	"golang-seed/apps/core/pkg/models"
	"golang-seed/pkg/server"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func registerRoutes(r *mux.Router) {
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := models.ConnectRepo(); err != nil {
		log.Fatal(err)
	}

	server := server.Init(config.Settings.Name, config.Settings.Port)
	server.ConfigureRouting()

	registerRoutes(server.RoutingRouter())

	server.Run()
}
