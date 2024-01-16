package cmd

import (
	"github.com/go-chi/chi/v5"
	"osrs-track-search/internal/handler"
)

type App struct {
	router  *chi.Mux
	handler *handler.Handler
}

func NewApp(router *chi.Mux, handler *handler.Handler) *App {
	return &App{router: router, handler: handler}
}

func (a *App) routes() {

	a.router.Get("/ironman", a.handler.SearchIronman)
}
