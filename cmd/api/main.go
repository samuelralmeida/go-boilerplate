package main

import (
	"flag"
	"log"

	"github.com/samuelralmeida/go-boilerplate/cmd/api/handlers"
	"github.com/samuelralmeida/go-boilerplate/internal/repository"
	"github.com/samuelralmeida/go-boilerplate/internal/server"
	"github.com/samuelralmeida/go-boilerplate/internal/service"
)

type config struct {
	addr string
	env  string
	tls  struct {
		certFile string
		keyFile  string
	}
}

type application struct {
	config  config
	handler handlers.Handler
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "localhost:4242", "server address to listen on")
	flag.StringVar(&cfg.env, "env", "development", "environment")

	repository := repository.New(repository.Options{})
	service := service.New(service.Options{Repository: repository})
	handlers := handlers.New(handlers.Options{Service: service})

	app := &application{
		config:  cfg,
		handler: handlers,
	}

	server := server.New(server.Options{Addr: cfg.addr, Routes: app.routes(), CertFile: cfg.tls.certFile, KeyFile: cfg.tls.keyFile})

	log.Printf("starting server on %s\n", cfg.addr)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
