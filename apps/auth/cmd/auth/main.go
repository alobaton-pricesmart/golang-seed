package main

import (
	"golang-seed/apps/auth/pkg/config"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service/auth"
	"golang-seed/apps/auth/pkg/store"
	"golang-seed/pkg/service"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func registerRoutes(r *mux.Router) {
	manager := manage.NewDefaultManager()
	// token store
	manager.MustTokenStorage(store.NewTokenStore())

	// client store
	manager.MapClientStorage(store.NewClientStore())

	// auth server
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	authService := auth.NewAuthService(srv)
	r.HandleFunc("/oauth/authorize", authService.Authorize)
	r.HandleFunc("/oauth/token", authService.Token)
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := models.ConnectRepo(); err != nil {
		log.Fatal(err)
	}

	service := service.Init("auth")
	service.ConfigureRouting()

	registerRoutes(service.RoutingRouter())

	service.Run()
}
