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
	"net/http"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauth2server "github.com/go-oauth2/oauth2/v4/server"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func registerRoutes(r *mux.Router) {
	// Set up your services first.
	clientsService := clientsserv.NewClientsService()
	usersService := usersserv.NewUsersService()

	// auth handler
	manager := manage.NewDefaultManager()
	// To setup the token duration user manager.SetAuthorizeCodeTokenCfg, manager.SetImplicitTokenCfg,
	// manager.SetPasswordTokenCfg, manager.SetClientTokenCfg, manager.SetRefreshTokenCfg
	// token store
	manager.MustTokenStorage(store.NewTokenStore())

	// client store
	manager.MapClientStorage(store.NewClientStore())

	// auth server
	srv := oauth2server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.ClientCredentials, oauth2.Refreshing)
	srv.SetClientInfoHandler(oauth2server.ClientBasicHandler)
	srv.SetAllowedResponseType(oauth2.Token)

	authHandler := authhand.NewAuthHandler(srv, usersService)
	r.HandleFunc("/oauth/authorize", authHandler.Authorize)
	r.HandleFunc("/oauth/token", authHandler.Token)

	srv.SetInternalErrorHandler(authHandler.InternalErrorHandler)
	srv.SetPasswordAuthorizationHandler(authHandler.PasswordAuthorizationHandler)

	// clients handler
	clientsHandler := clientshand.NewClientsHandler(clientsService)
	s := r.PathPrefix("/clients").Subrouter()
	s.Handle("/{id}", middleware.AuthenticationHandler(middleware.ErrorHandler(clientsHandler.Get), authHandler.ValidateToken)).Methods(http.MethodGet)
	s.Handle("/search/list", middleware.ErrorHandler(clientsHandler.GetAll)).Methods(http.MethodGet)
	s.Handle("/search/paged", middleware.ErrorHandler(clientsHandler.GetAllPaged)).Methods(http.MethodGet)
	s.Handle("/", middleware.ErrorHandler(clientsHandler.Create)).Methods(http.MethodPost)
	s.Handle("/{id}", middleware.ErrorHandler(clientsHandler.Update)).Methods(http.MethodPut)
	s.Handle("/{id}", middleware.ErrorHandler(clientsHandler.Delete)).Methods(http.MethodDelete)

	// users handler
	usersHandler := usershand.NewUsersHandler(usersService)
	s = r.PathPrefix("/users").Subrouter()
	s.Handle("/{id}", middleware.ErrorHandler(usersHandler.Get)).Methods(http.MethodGet)
	s.Handle("/search/list", middleware.ErrorHandler(usersHandler.GetAll)).Methods(http.MethodGet)
	s.Handle("/search/paged", middleware.ErrorHandler(usersHandler.GetAllPaged)).Methods(http.MethodGet)
	s.Handle("/", middleware.ErrorHandler(usersHandler.Create)).Methods(http.MethodPost)
	s.Handle("/{id}", middleware.ErrorHandler(usersHandler.Update)).Methods(http.MethodPut)
	s.Handle("/{id}", middleware.ErrorHandler(usersHandler.Delete)).Methods(http.MethodDelete)
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := models.ConnectRepo(); err != nil {
		log.Fatal(err)
	}

	if err := messages.Init("apps/auth/config", "es"); err != nil {
		log.Fatal(err)
	}

	server := server.Init(config.Settings.Name, config.Settings.Port)
	server.ConfigureRouting()

	registerRoutes(server.RoutingRouter())

	server.Run()
}
