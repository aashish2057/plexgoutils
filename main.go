package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aashish2057/plexgoutils/ombi_server"
	"github.com/aashish2057/plexgoutils/server"
	_ "github.com/aashish2057/plexgoutils/sessions"
)

func main() {
	godotenv.Load()
	var server1 server.Server
	var ombi server.Ombi
	server1.SetUrl(os.Getenv("PLEX_URL"))
	server1.SetToken(os.Getenv("PLEX_TOKEN"))

	ombi.SetUrl(os.Getenv("OMBI_URL"))
	ombi.SetKey(os.Getenv("OMBI_APIKEY"))

	// res, err := sessions.GetAccounts(server1)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	res1, err1 := ombi_server.GetAllMovieRequests(ombi)
	// res2, err2 := ombi_server.GetAllTvRequests(ombi)

	if err1 != nil {
		fmt.Println(err1)
		// fmt.Println(err2)
	} else {
		fmt.Println(res1)
		// fmt.Println(res2)
	}
}
