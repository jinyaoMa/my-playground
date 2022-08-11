package app

import (
	"my-playground/backend/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	http  *http.Server // redirector
	https *http.Server // server (tls)
}

func NewServer() *Server {
	handler := gin.Default()

	router.Setup(handler)

	s := &Server{
		http: &http.Server{},
		https: &http.Server{
			Addr:    ":10003",
			Handler: handler,
		},
	}

	go s.https.ListenAndServe()

	return s
}

func (s *Server) Start() *Server {
	return s
}

func (s *Server) Stop() *Server {
	return s
}
