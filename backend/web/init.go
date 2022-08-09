package web

import (
	"sync"
)

var once sync.Once
var web *Web

func Server() *Web {
	once.Do(func() {
		web = &Web{}
	})

	return web
}
