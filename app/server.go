package main

import (
	"log"

	"witzepedia/pkg/app"
	"witzepedia/pkg/config"
)

func main() {
	cfg := config.LoadConfig()
	server := app.NewServer(cfg)

	log.Printf("Starting HTTP Server. Listening at %q", cfg.Addr)
	log.Fatal(server.ListenAndServe())
}
