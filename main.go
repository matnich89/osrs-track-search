package main

import (
	"github.com/go-chi/chi/v5"
	cmd "github.com/matnich89/osrs-track-search/cmd/api"
	"github.com/matnich89/osrs-track-search/internal/client"
	"github.com/matnich89/osrs-track-search/internal/handler"
	"log"
	"os"
)

func main() {

	router := chi.NewMux()

	jagexHost, ok := os.LookupEnv("JAGEX_HOST")

	if !ok {
		log.Fatal("could not find jagex host url")
	}

	c, err := client.NewJagexClient(jagexHost)

	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(c)

	app := cmd.NewApp(router, h)

	err = app.ConnectToNats()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Serve()

	if err != nil {
		log.Fatal(err)
	}

}
