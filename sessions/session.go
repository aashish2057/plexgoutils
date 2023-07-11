package sessions

import (
	"fmt"

	"github.com/aashish2057/plexgoutils/server"
)

func GetActiveSessions(s server.Server) {
	fmt.Println("Url: " + s.url + " Token: " + s.token)
}
