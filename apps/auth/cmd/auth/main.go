package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/go-oauth2/mysql.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"

	"github.com/alobaton/golang-seed/apps/auth/pkg/config"
)

var (
	DnsChain = "%s:%s@%s?charset=utf8"
)

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	manager := manage.NewDefaultManager()

	dsn := fmt.Sprintf(DnsChain, config.Database.User, config.Database.Password, config.Database.Address)
	store := mysql.NewDefaultStore(
		mysql.NewConfig(dsn),
	)
	defer store.Close()

	manager.MapTokenStorage(store)
	manager.MapClientStorage(store)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("internal error: ", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("response error: ", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
