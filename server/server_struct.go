package server

import (
	"net/http"
	"github.com/tuyensinh/utils"
)

type Server struct {
	config               utils.ConfigServer
	mux                  map[string]func(http.ResponseWriter, *http.Request)
	server               *http.Server
}
