package settings

import (
	"net/http"
)

type SettingsServer struct{}

func NewSettingsServer() *SettingsServer {
	return new(SettingsServer)
}

func (s *SettingsServer) CreateSetting(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
