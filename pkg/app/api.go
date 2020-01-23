package app

import (
	"net/http"

	"witzepedia/pkg/auth"
	"witzepedia/pkg/config"
)

func NewServer(cfg *config.Config) *http.Server {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	auth.MountHandlers(mux)

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}
	return server
}
