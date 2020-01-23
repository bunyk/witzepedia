package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Addr  string
	OAuth oauth2.Config
}

var config *Config

func LoadConfig() *Config {
	if config == nil {
		config = &Config{
			Addr: env("APP_PORT", ":8001"),
			OAuth: oauth2.Config{
				RedirectURL:  "http://localhost:8000/auth/google/callback",
				ClientID:     env("GOOGLE_OAUTH_CLIENT_ID", "id"),
				ClientSecret: env("GOOGLE_OAUTH_CLIENT_SECRET", "secret"),
				Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
				Endpoint:     google.Endpoint,
			},
		}
	}
	return config
}

func env(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		val = def
		log.Printf("Env variable %s not set, using default value \"%s\"", name, val)
	}
	return val
}
