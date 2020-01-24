package config

import (
	"fmt"
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
		port := env("APP_PORT", "8001")
		domain := env("APP_DOMAIN", "witzepedia.org")
		config = &Config{
			Addr: ":" + port,
			OAuth: oauth2.Config{
				RedirectURL:  fmt.Sprintf("http://%s:%s/auth/google/callback", domain, port),
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
