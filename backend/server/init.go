package server

import (
	"sync"
)

var once sync.Once
var server *Server

func One() *Server {
	once.Do(func() {
		server = &Server{}
	})

	return server
}
