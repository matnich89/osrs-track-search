package cmd

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/matnich89/osrs-track-search/internal/handler"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	router  *chi.Mux
	handler *handler.Handler
}

func NewApp(router *chi.Mux, handler *handler.Handler) *App {
	return &App{router: router, handler: handler}
}

func (a *App) ConnectToNats() error {

	url, ok := os.LookupEnv("NATS_URL")

	if !ok {
		return errors.New("could not find NATS_URL env")
	}

	nc, err := nats.Connect(url)

	if err != nil {
		return err
	}

	log.Println("connected to nats......")

	defer nc.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Closing connection to NATS")

	return nil

}

func (a *App) routes() {
	a.router.Get("/ironman", a.handler.SearchIronman)
}
