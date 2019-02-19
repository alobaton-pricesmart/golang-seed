package masters

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MastersServer struct{}

func NewMastersSever() *MastersServer {
	return new(MastersServer)
}

func (s *MastersServer) CreateMaster(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
