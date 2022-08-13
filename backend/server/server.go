package server

import (
	"context"
	"crypto/tls"
	"embed"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

//go:embed certs
var certs embed.FS

var server *Server

type Server struct {
	http        *http.Server // redirector
	https       *http.Server // server (tls)
	errGroup    errgroup.Group
	httpPort    string
	httpsPort   string
	certManager *autocert.Manager
	tlsConfig   *tls.Config
}

func Setup() {
	port := ":10080"
	portTls := ":10443"
	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  nil,
	}
	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = server.getSelfSignedOrLetsEncryptCert(manager)

	server = &Server{
		httpPort:    port,
		httpsPort:   portTls,
		certManager: manager,
		tlsConfig:   tlsConfig,
	}
}

func StartServer() (ok bool) {
	server.http = &http.Server{
		Addr: server.httpPort,
		Handler: server.certManager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, server.httpPort, server.httpsPort, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
		ErrorLog: nil,
	}
	server.https = &http.Server{
		Addr:      server.httpsPort,
		TLSConfig: server.tlsConfig,
		Handler:   setup(gin.Default()),
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

	if err := server.http.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Server (HTTP) shutdown error: %+v\n", err)
	}
	if err := server.https.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Server (HTTP/S) shutdown error: %+v\n", err)
	}

	return true
}

func (s *Server) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		var certificate tls.Certificate
		var err error
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if ok {
			keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
			crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
			certificate, err = tls.LoadX509KeyPair(crtFile, keyFile)
		} else {
			key, _ := certs.ReadFile("certs/localhost.key")
			crt, _ := certs.ReadFile("certs/localhost.crt")
			certificate, err = tls.X509KeyPair(crt, key)
		}
		if err != nil {
			log.Printf("%s\nFalling back to Letsencrypt\n", err)
			return certManager.GetCertificate(hello)
		}
		log.Println("Loaded selfsigned certificate.")
		return &certificate, err
	}
}
