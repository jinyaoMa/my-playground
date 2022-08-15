package server

import (
	"context"
	"crypto/tls"
	"embed"
	"log"
	"my-playground/backend/utils"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

//go:embed certs
var certs embed.FS

var server *Server

func init() {
	server = &Server{}
}

type Server struct {
	isRunning    bool
	mtx          sync.Mutex
	errGroup     errgroup.Group
	http         *http.Server // redirector
	https        *http.Server // server (tls)
	tlsConfig    *tls.Config
	httpHandler  http.Handler
	httpsHandler http.Handler
	config       *Config
}

const (
	CfgNameHttpPort      = "Server.HttpPort"
	CfgNameHttpsPort     = "Server.HttpsPort"
	CfgNameCertsDirCache = "Server.CertsDirCache"
)

type Config struct {
	HttpPort      string
	HttpsPort     string
	CertsDirCache string
}

func SetConfig(cfg *Config) {
	Stop() // stop server if running

	server.config = cfg

	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
	}
	if cfg.CertsDirCache == "" {
		manager.Cache = nil
	} else if utils.IsDirectoryExist(cfg.CertsDirCache) {
		manager.Cache = autocert.DirCache(cfg.CertsDirCache)
	}
	tlsConfig := manager.TLSConfig()
	tlsConfig.GetCertificate = server.getSelfSignedOrLetsEncryptCert(manager)

	server = &Server{
		tlsConfig: tlsConfig,
		httpHandler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, cfg.HttpPort, cfg.HttpsPort, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
		httpsHandler: setup(gin.Default()),
	}
}

func Start() (ok bool) {
	server.mtx.Lock()
	defer server.mtx.Unlock()
	if server.isRunning { // already running, cannot start again
		return false
	}

	server.http = &http.Server{
		Addr:     server.config.HttpPort,
		Handler:  server.httpHandler,
		ErrorLog: nil,
	}
	server.https = &http.Server{
		Addr:      server.config.HttpsPort,
		Handler:   server.httpsHandler,
		TLSConfig: server.tlsConfig,
		ErrorLog:  nil,
	}

	server.errGroup.Go(func() error {
		return server.http.ListenAndServe()
	})
	server.errGroup.Go(func() error {
		return server.https.ListenAndServeTLS("", "")
	})

	server.isRunning = true
	return true
}

func Stop() (ok bool) {
	server.mtx.Lock()
	defer server.mtx.Unlock()
	if !server.isRunning { // already stop or not yet start
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.http.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Server (HTTP) shutdown error: %+v\n", err)
	}
	if err := server.https.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Server (HTTP/S) shutdown error: %+v\n", err)
	}

	if err := server.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
		log.Printf("Server running error: %+v\n", err)
	}

	server.isRunning = false
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
