package server

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

var server *Server

type Server struct {
	http         *http.Server // redirector
	https        *http.Server // server (tls)
	errGroup     errgroup.Group
	httpPort     string
	httpsPort    string
	tlsConfig    *tls.Config
	httpHandler  http.Handler
	httpsHandler *gin.Engine
}

func Setup() {
	port := ":10080"
	portTls := ":10443"
	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("D:\\GitHub\\my-playground\\build\\bin"),
	}
	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = server.getSelfSignedOrLetsEncryptCert(manager)

	server = &Server{
		httpPort:  port,
		httpsPort: portTls,
		tlsConfig: tlsConfig,
		httpHandler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, port, portTls, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
		httpsHandler: setup(gin.Default()),
	}
}

func StartServer() (ok bool) {
	server.http = &http.Server{
		Addr:     server.httpPort,
		Handler:  server.httpHandler,
		ErrorLog: nil,
	}
	server.https = &http.Server{
		Addr:      server.httpsPort,
		TLSConfig: server.tlsConfig,
		Handler:   server.httpsHandler,
		ErrorLog:  nil,
	}

	server.errGroup.Go(func() error {
		return server.http.ListenAndServe()
	})
	server.errGroup.Go(func() error {
		return server.https.ListenAndServeTLS("", "")
	})

	return true
}

func StopServer() (ok bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.http.Shutdown(ctx)
	err = server.https.Shutdown(ctx)
	return err == nil || err == http.ErrServerClosed
}

func (s *Server) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if !ok {
			dirCache = "certs"
		}

		keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
		crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			log.Printf("%s\nFalling back to Letsencrypt\n", err)
			return certManager.GetCertificate(hello)
		}
		log.Println("Loaded selfsigned certificate.")
		return &certificate, err
	}
}
