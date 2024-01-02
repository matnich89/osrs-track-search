package main

import "github.com/go-chi/chi/v5"

type app struct {
	router *chi.Mux
}

func newApp(router *chi.Mux) *app {
	return &app{router: router}
}

func (a *app) routes() {
	a.router.Get("/ironman", a.han)
}
