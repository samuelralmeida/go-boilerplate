package server

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Options struct {
	Addr     string
	Routes   http.Handler
	CertFile string
	KeyFile  string
}

type server struct {
	srv      *http.Server
	certFile string
	keyFile  string
}

func New(options Options) *server {
	srv := newServer(options.Addr, options.Routes)
	return &server{
		srv:      srv,
		certFile: options.CertFile,
		keyFile:  options.KeyFile,
	}
}

func newServer(addr string, h http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      h,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *server) Run() error {
	shutdownError := s.makeShutdownErrorChan()

	var err error
	if s.hasTlsFiles() {
		err = s.runWithTls()
	} else {
		err = s.runWithoutTls()
	}

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return <-shutdownError
}

func (s *server) makeShutdownErrorChan() chan error {
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- s.srv.Shutdown(ctx)
	}()

	return shutdownError
}

func (s *server) hasTlsFiles() bool {
	return s.certFile != "" && s.keyFile != ""
}

func (s *server) runWithTls() error {
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	s.srv.TLSConfig = tlsConfig
	return s.srv.ListenAndServeTLS(s.certFile, s.keyFile)
}

func (s *server) runWithoutTls() error {
	return s.srv.ListenAndServe()
}
