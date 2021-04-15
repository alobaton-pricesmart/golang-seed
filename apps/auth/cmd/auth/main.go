package main

import (
	"golang-seed/apps/auth/pkg/config"
	"golang-seed/apps/auth/pkg/handler/authhand"
	"golang-seed/apps/auth/pkg/handler/clientshand"
	"golang-seed/apps/auth/pkg/handler/usershand"
	"golang-seed/apps/auth/pkg/models"
	"golang-seed/apps/auth/pkg/service/clientsserv"
	"golang-seed/apps/auth/pkg/service/usersserv"
	"golang-seed/apps/auth/pkg/store"
	"golang-seed/pkg/messages"
	"golang-seed/pkg/middleware"
	"golang-seed/pkg/server"

	"github.com/go-oauth2/oauth2/v4/manage"
	oauth2server "github.com/go-oauth2/oauth2/v4/server"
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
	srv := oauth2server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(oauth2server.ClientFormHandler)

	authHandler := authhand.NewAuthHandler(srv)
	r.HandleFunc("/oauth/authorize", authHandler.Authorize)
	r.HandleFunc("/oauth/token", authHandler.Token)

	usersService := usersserv.NewUsersService()
	usersHandler := usershand.NewUsersHandler(usersService)
	r.Handle("/users/{id}", middleware.ErrorHandler(usersHandler.Get)).Methods("GET")
	r.Handle("/users/search/list", middleware.ErrorHandler(usersHandler.GetAll)).Methods("GET")
	r.Handle("/users/search/paged", middleware.ErrorHandler(usersHandler.GetAllPaged)).Methods("GET")
	r.Handle("/users", middleware.ErrorHandler(usersHandler.Create)).Methods("POST")
	r.Handle("/users/{id}", middleware.ErrorHandler(usersHandler.Update)).Methods("PUT")
	r.Handle("/users/{id}", middleware.ErrorHandler(usersHandler.Delete)).Methods("DELETE")

	clientsService := clientsserv.NewClientsService()
	clientsHandler := clientshand.NewClientsHandler(clientsService)
	r = r.PathPrefix("/clients").Subrouter()
	r.Handle("/{id}", middleware.ErrorHandler(clientsHandler.Get)).Methods("GET")
	r.Handle("/search/list", middleware.ErrorHandler(clientsHandler.GetAll)).Methods("GET")
	r.Handle("/search/paged", middleware.ErrorHandler(clientsHandler.GetAllPaged)).Methods("GET")
	r.Handle("/", middleware.ErrorHandler(clientsHandler.Create)).Methods("POST")
	r.Handle("/{id}", middleware.ErrorHandler(clientsHandler.Update)).Methods("PUT")
	r.Handle("/{id}", middleware.ErrorHandler(clientsHandler.Delete)).Methods("DELETE")
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := models.ConnectRepo(); err != nil {
		log.Fatal(err)
	}

	if err := messages.Init("apps/auth/config/es.json", "es"); err != nil {
		log.Fatal(err)
	}

	server := server.Init(config.Settings.Name, config.Settings.Port)
	server.ConfigureRouting()

	registerRoutes(server.RoutingRouter())

	server.Run()
}
