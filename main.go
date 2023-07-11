package main

import (
	"fmt"

	"github.com/aashish2057/plexgoutils/server"
)

func main() {

	var s server.Server
	s.SetUrl("test url")
	s.SetToken("test token")
	fmt.Println(s.GetUrl())
	fmt.Println(s.GetToken())
}
