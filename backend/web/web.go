package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Web struct {
	isRunning bool
	http      *http.Server // redirecter
	https     *http.Server // server (tls)
}

func (s *Web) Start() {
	if s.isRunning {
		return
	}

	s.isRunning = true

	router := gin.Default()

	Setup(router)

	s.https = &http.Server{
		Addr:    ":10003",
		Handler: router,
	}
}

func (s *Web) Stop() {
	if !s.isRunning {
		return
	}

	s.isRunning = false
}
