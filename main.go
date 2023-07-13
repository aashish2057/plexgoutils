package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aashish2057/plexgoutils/server"
	"github.com/aashish2057/plexgoutils/sessions"
)

func main() {
	godotenv.Load()
	var server1 server.Server
	server1.SetUrl(os.Getenv("PLEX_URL"))
	server1.SetToken(os.Getenv("PLEX_TOKEN"))

	res, err := sessions.StopTranscodeSession(server1, "")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
