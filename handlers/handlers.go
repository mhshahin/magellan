package handlers

import "github.com/mhshahin/magellan/repository"

type Handlers struct {
	Router *RouterHandlers
}

func NewHandlers(repos *repository.Repository) *Handlers {
	return &Handlers{
		Router: NewRouterHandlers(repos),
	}
}
