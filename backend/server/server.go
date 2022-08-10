package server

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once   sync.Once
	server *Server
)

type Server struct {
	redirector *http.Server // redirector
	server     *http.Server // server (tls)
}

func Load() *Server {
	once.Do(func() {
		router := gin.Default()

		setup(router)

		server = &Server{
			redirector: &http.Server{},
			server: &http.Server{
				Addr:    ":10003",
				Handler: router,
			},
		}
	})
	return server
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

// run on air
func Main() {
	main()
}

// run on air
func main() {
	Load()
}
