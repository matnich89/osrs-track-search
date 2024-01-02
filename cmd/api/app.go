package cmd

import (
	"github.com/go-chi/chi/v5"
	"osrs-track-search/internal/handler"
)

type app struct {
	router  *chi.Mux
	handler *handler.Handler
}

func NewApp(router *chi.Mux) *app {
	return &app{router: router}
}

func (a *app) routes() {
	a.router.Get("/ironman", a.handler.SearchIronman)
}
