package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	cmd "osrs-track-search/cmd/api"
	"osrs-track-search/internal/client"
	"osrs-track-search/internal/handler"
)

func main() {

	router := chi.NewMux()

	c, err := client.NewJagexClient("https://secure.runescape.com")

	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(c)

	app := cmd.NewApp(router, h)

	err = app.Serve()

	if err != nil {
		log.Fatal(err)
	}

}
