# Plexgoutils

plexgoutils is a collection of go packages to perform plex server functionality

## Packages

### Server

This package contains the struct that store your plex server url and token. This struct is used in almost all of the functions and is essential to using these packages.

How to setup the server struct, it it recommended that you use a .env file and load it in for your url and tokens however that is not required. Initialize a variable of type Server and call the SetUrl and SetToken functions. To validate everything is correct you can call GetUrl and GetToken to see values the struct is holding.

    package main

    import (
    	"github.com/aashish2057/plexgoutils/server"
    )

    func main() {
    	var server1 server.Server
    	server1.SetUrl("PLEX_URL")
    	server1.SetToken("PLEX_TOKEN")

    	server1.GetUrl()
    	server1.GetToken()
     }

### Sessions

This package contains all the functions relating to sessions on a plex server.

#### GetActiveSessions() -> returns all active sessions on the plex server

#### GetSessionHistory() -> returns all active sessions on the plex server

#### GetTranscodeSessions() -> returns all active transcode sessions on the plex server

#### StopTranscodeSession() -> stops specific transcode session

All of these functions require the server struct from the server package. For Stop Transcode Session you also need to provide the transcode session key, which can be obtained from the GetTranscodeSessions() response

    package main

    import (
    	"fmt"

    	"github.com/aashish2057/plexgoutils/server"
    	"github.com/aashish2057/plexgoutils/sessions"
    )

    func main() {
    	godotenv.Load()
    	var server1 server.Server
    	server1.SetUrl("PLEX_URL")
    	server1.SetToken("PLEX_TOKEN")

    	res, err := sessions.GetActiveSessions(server1)
    	res, err := sessions.GetSessionHistory(server1)
    	res, err := sessions.GetTranscodeSessions(server1)
    	res, err := sessions.StopTranscodeSession(server1, "session key")
    }
