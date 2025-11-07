package main

import (
	"log"

	"oilio.com/internal/config"
	httpserver "oilio.com/internal/http"
)

func main() {
	cfg := config.Load()
	app := httpserver.New()
	log.Printf("Listening on %s", cfg.ServerPort)
	if err := app.Listen(cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
