package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	isRunning bool
	http      *http.Server // redirecter
	https     *http.Server // server (tls)
}

func (s *Server) Start(swag bool, loop bool) (ok bool) {
	if s.isRunning {
		return false
	}

	s.isRunning = true
	s.https = &http.Server{
		Addr:    ":10003",
		Handler: gin.Default(),
	}

	return true
}

func (s *Server) Stop() (ok bool) {
	if !s.isRunning {
		return false
	}

	s.isRunning = false
	return true
}
