package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	cmd "osrs-track-search/cmd/api"
)

func main() {

	router := chi.NewMux()

	app := cmd.NewApp(router)

	err := app.Serve()

	if err != nil {
		log.Fatal(err)
	}

}
