package server

import (
	"context"
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"my-playground/backend/utils"
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

var (
	server *Server
	config *Config
)

type Server struct {
	errGroup     errgroup.Group
	http         *http.Server // redirector
	https        *http.Server // server (tls)
	httpPort     string
	httpsPort    string
	tlsConfig    *tls.Config
	httpHandler  http.Handler
	httpsHandler http.Handler
}

type Config struct {
	HttpPort      string
	HttpsPort     string
	CertsDirCache string
}

func Setup(cfg *Config) {
	config = cfg

	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
	}
	if config.CertsDirCache == "" {
		manager.Cache = nil
	} else if utils.IsDirectoryExist(config.CertsDirCache) {
		manager.Cache = autocert.DirCache(config.CertsDirCache)
	}
	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = server.getSelfSignedOrLetsEncryptCert(manager)

	server = &Server{
		httpPort:  config.HttpPort,
		httpsPort: config.HttpsPort,
		tlsConfig: tlsConfig,
		httpHandler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, config.HttpPort, config.HttpsPort, 1) + r.RequestURI
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
		Handler:   server.httpsHandler,
		TLSConfig: server.tlsConfig,
		ErrorLog:  nil,
	}
	fmt.Printf("%s\n", server.httpPort)
	fmt.Printf("%s\n", server.httpsPort)

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
