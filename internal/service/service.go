package service

import (
	"fmt"

	"github.com/samuelralmeida/go-boilerplate/internal/repository"
)

type Servicer interface {
	Home() error
}

type service struct {
	repository repository.Repositorier
}

type Options struct {
	Repository repository.Repositorier
}

func New(options Options) Servicer {
	service := new(service)
	service.repository = options.Repository
	return service
}

func (h *service) Home() error {
	fmt.Println("service calling repository")
	return h.repository.Home()
}
