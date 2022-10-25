//go:generate mockery --output ../mocks --name Repositorier

package repository

import "fmt"

type Repositorier interface {
	Home() error
}

type repository struct {
	// db database.DB
}

type Options struct {
	// DB database.DB
}

func New(options Options) Repositorier {
	repository := new(repository)
	// handler.service = options.Service
	return repository
}

func (h *repository) Home() error {
	fmt.Println("welcome to repository")
	return nil
}
