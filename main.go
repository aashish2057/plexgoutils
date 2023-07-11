package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aashish2057/plexgoutils/server"
)

func main() {
	godotenv.Load()
	var s server.Server
	s.SetUrl(os.Getenv("PLEX_URL"))
	s.SetToken(os.Getenv("PLEX_TOKEN"))
}
