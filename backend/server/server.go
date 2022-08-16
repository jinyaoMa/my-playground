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

const (
	PkgName = "server"

	CfgNameHttpPort      = "Server.HttpPort"
	CfgNameHttpsPort     = "Server.HttpsPort"
	CfgNameCertsDirCache = "Server.CertsDirCache"
)

var (
	server *Server
)

func init() {
	server = &Server{}
	utils.Logger(PkgName).Println("SERVER INIT")
}

type Config struct {
	HttpPort      string
	HttpsPort     string
	CertsDirCache string
}

type Server struct {
	isRunning    bool
	config       *Config
	mtx          sync.Mutex
	errGroup     errgroup.Group
	tlsConfig    *tls.Config
	httpHandler  http.Handler
	httpsHandler http.Handler
	http         *http.Server // redirector
	https        *http.Server // server (tls)
}

func My() *Server {
	return server
}

func (s *Server) SetConfig(cfg *Config) *Server {
	needRestart := s.Stop() // stop server if running

	s.mtx.Lock()
	{
		s.config = cfg

		manager := &autocert.Manager{
			Prompt: autocert.AcceptTOS,
		}
		if cfg.CertsDirCache == "" {
			manager.Cache = nil
		} else if utils.IsDirectoryExist(cfg.CertsDirCache) {
			manager.Cache = autocert.DirCache(cfg.CertsDirCache)
		}

		s.tlsConfig = manager.TLSConfig()
		s.tlsConfig.GetCertificate = s.getSelfSignedOrLetsEncryptCert(manager)

		s.httpHandler = manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, cfg.HttpPort, cfg.HttpsPort, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		}))

		s.httpsHandler = SetupHandler(gin.Default())
	}
	s.mtx.Unlock()

	if needRestart { // restart server if it's stopped to set config
		s.Start()
	}
	return s
}

func (s *Server) Start() (ok bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.isRunning { // already running, cannot start again
		return false
	}

	s.http = &http.Server{
		Addr:     s.config.HttpPort,
		Handler:  s.httpHandler,
		ErrorLog: nil,
	}
	s.https = &http.Server{
		Addr:      s.config.HttpsPort,
		Handler:   s.httpsHandler,
		TLSConfig: s.tlsConfig,
		ErrorLog:  nil,
	}

	s.errGroup.Go(func() error {
		return s.http.ListenAndServe()
	})
	s.errGroup.Go(func() error {
		return s.https.ListenAndServeTLS("", "")
	})

	s.isRunning = true
	utils.Logger(PkgName).Println("SERVER START")
	return true
}

func (s *Server) Stop() (ok bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if !s.isRunning { // already stop or not yet start
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Server (HTTP) shutdown error: %+v\n", err)
	}
	if err := s.https.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Server (HTTP/S) shutdown error: %+v\n", err)
	}

	if err := s.errGroup.Wait(); err != nil && err != http.ErrServerClosed {
		log.Printf("Server running error: %+v\n", err)
	}

	s.isRunning = false
	utils.Logger(PkgName).Println("SERVER STOP")
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
