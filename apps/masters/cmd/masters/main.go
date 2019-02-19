package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/alobaton/golang-seed/apps/masters/pkg/config"
	"github.com/alobaton/golang-seed/apps/masters/pkg/models"
	"github.com/alobaton/golang-seed/apps/masters/pkg/services/masters"
	"github.com/alobaton/golang-seed/pkg/services"
)

func registerRoutes(r *httprouter.Router) {
	server := masters.NewMastersSever()

	r.POST("/", server.CreateMaster)
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := models.ConnectRepo(); err != nil {
		log.Fatal(err)
	}

	service := services.Init("masters")
	service.ConfigureRouting()

	registerRoutes(service.RoutingRouter())

	service.Run()
}
