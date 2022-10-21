package main

import (
	"flag"
	"log"

	"github.com/samuelralmeida/go-boilerplate/internal/server"
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
	config config
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "localhost:4242", "server address to listen on")
	flag.StringVar(&cfg.env, "env", "development", "environment")

	app := &application{
		config: cfg,
	}

	handlers := app.routes()

	server := server.New(server.Options{Addr: cfg.addr, Handler: handlers, CertFile: cfg.tls.certFile, KeyFile: cfg.tls.keyFile})

	log.Printf("starting server on %s\n", cfg.addr)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
