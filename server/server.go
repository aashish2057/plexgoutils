package server

type Server struct {
	url   string
	token string
}

func (server *Server) SetUrl(url string) {
	server.url = url
}

func (server *Server) SetToken(token string) {
	server.token = token
}

func (server *Server) GetUrl() string {
	return server.url
}

func (server *Server) GetToken() string {
	return server.token
}
