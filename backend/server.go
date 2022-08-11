package backend

import (
	"my-playground/backend/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	redirector *http.Server // redirector
	server     *http.Server // server (tls)
)

func setupServer() {
	handler := gin.Default()

	router.Setup(handler)

	redirector = &http.Server{}
	server = &http.Server{
		Addr:    ":10003",
		Handler: handler,
	}
}

func Quit() {

}
