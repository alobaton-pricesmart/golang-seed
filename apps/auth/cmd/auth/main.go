package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-oauth2/mysql.v3"
	oa2errors "gopkg.in/oauth2.v3/errors"
	oa2manage "gopkg.in/oauth2.v3/manage"
	oa2server "gopkg.in/oauth2.v3/server"
	"github.com/julienschmidt/httprouter"

	"github.com/alobaton/golang-seed/apps/auth/pkg/config"
	"github.com/alobaton/golang-seed/apps/auth/pkg/store"
	"github.com/alobaton/golang-seed/pkg/services"
)

var (
	DnsChain = "%s:%s@tcp(%s)/app_auth?charset=utf8mb4"
)

func registerRoutes(r *httprouter.Router, s *oa2server.Server) {
	r.HandlerFunc("POST", "/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := s.HandleAuthorizeRequest(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	r.HandlerFunc("POST", "/token", func(w http.ResponseWriter, r *http.Request) {
		s.HandleTokenRequest(w, r)
	})
}

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	manager := oa2manage.NewDefaultManager()

	dns := fmt.Sprintf(DnsChain, config.Settings.Database.User, config.Settings.Database.Password, config.Settings.Database.Address)
	tokenStore := mysql.NewDefaultStore(
		mysql.NewConfig(dns),
	)
	defer tokenStore.Close()

	log.WithField("dns", dns).Debug("token store created successfully!")
	manager.MapTokenStorage(tokenStore)

	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	server := oa2server.NewDefaultServer(manager)
	server.SetAllowGetAccessRequest(true)
	server.SetClientInfoHandler(oa2server.ClientFormHandler)

	server.SetInternalErrorHandler(func(err error) (re *oa2errors.Response) {
		log.Error("internal error: ", err.Error())
		return
	})

	server.SetResponseErrorHandler(func(re *oa2errors.Response) {
		log.Error("response error: ", re.Error.Error())
	})

	service := services.Init("auth")
	service.ConfigureRouting()

	registerRoutes(service.RoutingRouter(), server)

	service.Run()
}
