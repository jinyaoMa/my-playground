package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	redirector *http.Server // redirector
	server     *http.Server // server (tls)
)

func Setup() {
	handler := gin.Default()

	setup(handler)

	redirector = &http.Server{}
	server = &http.Server{
		Addr:    ":10003",
		Handler: handler,
	}
}
