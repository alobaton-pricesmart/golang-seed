package main

import (
	"net/http"

	"golang-seed/apps/auth/pkg/config"
	"golang-seed/apps/auth/pkg/handler"
	"golang-seed/apps/auth/pkg/repo"
	"golang-seed/apps/auth/pkg/service"
	"golang-seed/apps/auth/pkg/store"
	"golang-seed/pkg/messages"
	"golang-seed/pkg/server"
	"golang-seed/pkg/server/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauth2server "github.com/go-oauth2/oauth2/v4/server"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := repo.ConnectRepo(); err != nil {
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

func registerRoutes(r *mux.Router) {
	// Set up your services first.
	clientsService := service.NewClientsService()
	usersService := service.NewUsersService()
	permissionsService := service.NewPermissionsService()
	rolesService := service.NewRolesService()

	// auth handler
	manager := manage.NewDefaultManager()
	// To setup the token duration user manager.SetAuthorizeCodeTokenCfg, manager.SetImplicitTokenCfg,
	// manager.SetPasswordTokenCfg, manager.SetClientTokenCfg, manager.SetRefreshTokenCfg
	// token store
	manager.MustTokenStorage(store.NewTokenStore())
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(config.Settings.Security.KeyID,
		[]byte(config.Settings.Security.Key),
		jwt.SigningMethodHS512))

	// client store
	manager.MapClientStorage(store.NewClientStore())

	// auth server
	srv := oauth2server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.ClientCredentials, oauth2.Refreshing)
	srv.SetClientInfoHandler(oauth2server.ClientBasicHandler)
	srv.SetAllowedResponseType(oauth2.Token)

	authHandler := handler.NewAuthHandler(srv, usersService)
	r.HandleFunc("/oauth/authorize", authHandler.Authorize)
	r.HandleFunc("/oauth/token", authHandler.Token)

	srv.SetInternalErrorHandler(authHandler.InternalErrorHandler)
	srv.SetPasswordAuthorizationHandler(authHandler.PasswordAuthorizationHandler)

	// clients handler
	clientsHandler := handler.NewClientsHandler(clientsService)
	s := r.PathPrefix("/clients").Subrouter()
	s.Use(middleware.AuthenticationHandler(authHandler.ValidateToken))
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(clientsHandler.Get),
		middleware.AuthorizeHandler("read:client", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/list", middleware.Middleware(
		middleware.ErrorHandler(clientsHandler.GetAll),
		middleware.AuthorizeHandler("read:clients", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/paged", middleware.Middleware(
		middleware.ErrorHandler(clientsHandler.GetAllPaged),
		middleware.AuthorizeHandler("read:clients", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("", middleware.Middleware(
		middleware.ErrorHandler(clientsHandler.Create),
		middleware.AuthorizeHandler("create:client", authHandler.ValidatePermission)),
	).Methods(http.MethodPost)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(clientsHandler.Update),
		middleware.AuthorizeHandler("update:client", authHandler.ValidatePermission)),
	).Methods(http.MethodPut)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(clientsHandler.Delete),
		middleware.AuthorizeHandler("delete:client", authHandler.ValidatePermission)),
	).Methods(http.MethodDelete)

	// users handler
	usersHandler := handler.NewUsersHandler(usersService)
	s = r.PathPrefix("/users").Subrouter()
	s.Use(middleware.AuthenticationHandler(authHandler.ValidateToken))
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(usersHandler.Get),
		middleware.AuthorizeHandler("get:user", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/list", middleware.Middleware(
		middleware.ErrorHandler(usersHandler.GetAll),
		middleware.AuthorizeHandler("get:users", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/paged", middleware.Middleware(
		middleware.ErrorHandler(usersHandler.GetAllPaged),
		middleware.AuthorizeHandler("get:users", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("", middleware.Middleware(
		middleware.ErrorHandler(usersHandler.Create),
		middleware.AuthorizeHandler("create:user", authHandler.ValidatePermission)),
	).Methods(http.MethodPost)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(usersHandler.Update),
		middleware.AuthorizeHandler("update:user", authHandler.ValidatePermission)),
	).Methods(http.MethodPut)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(usersHandler.Delete),
		middleware.AuthorizeHandler("delete:user", authHandler.ValidatePermission)),
	).Methods(http.MethodDelete)

	// permissions handler
	permissionsHandler := handler.NewPermissionsHandler(permissionsService)
	s = r.PathPrefix("/permissions").Subrouter()
	s.Use(middleware.AuthenticationHandler(authHandler.ValidateToken))
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(permissionsHandler.Get),
		middleware.AuthorizeHandler("get:permission", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/list", middleware.Middleware(
		middleware.ErrorHandler(permissionsHandler.GetAll),
		middleware.AuthorizeHandler("get:permissions", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/paged", middleware.Middleware(
		middleware.ErrorHandler(permissionsHandler.GetAllPaged),
		middleware.AuthorizeHandler("get:permissions", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("", middleware.Middleware(
		middleware.ErrorHandler(permissionsHandler.Create),
		middleware.AuthorizeHandler("create:permission", authHandler.ValidatePermission)),
	).Methods(http.MethodPost)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(permissionsHandler.Update),
		middleware.AuthorizeHandler("update:permission", authHandler.ValidatePermission)),
	).Methods(http.MethodPut)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(permissionsHandler.Delete),
		middleware.AuthorizeHandler("delete:permission", authHandler.ValidatePermission)),
	).Methods(http.MethodDelete)

	// roles handler
	rolesHandler := handler.NewRolesHandler(rolesService)
	s = r.PathPrefix("/roles").Subrouter()
	s.Use(middleware.AuthenticationHandler(authHandler.ValidateToken))
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(rolesHandler.Get),
		middleware.AuthorizeHandler("get:role", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/list", middleware.Middleware(
		middleware.ErrorHandler(rolesHandler.GetAll),
		middleware.AuthorizeHandler("get:roles", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("/search/paged", middleware.Middleware(
		middleware.ErrorHandler(rolesHandler.GetAllPaged),
		middleware.AuthorizeHandler("get:roles", authHandler.ValidatePermission)),
	).Methods(http.MethodGet)
	s.Handle("", middleware.Middleware(
		middleware.ErrorHandler(rolesHandler.Create),
		middleware.AuthorizeHandler("create:role", authHandler.ValidatePermission)),
	).Methods(http.MethodPost)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(rolesHandler.Update),
		middleware.AuthorizeHandler("update:role", authHandler.ValidatePermission)),
	).Methods(http.MethodPut)
	s.Handle("/{id}", middleware.Middleware(
		middleware.ErrorHandler(rolesHandler.Delete),
		middleware.AuthorizeHandler("delete:role", authHandler.ValidatePermission)),
	).Methods(http.MethodDelete)
}
