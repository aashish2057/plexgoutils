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

type Ombi struct {
	url    string
	apikey string
}

func (ombi *Ombi) SetUrl(url string) {
	ombi.url = url
}

func (ombi *Ombi) SetKey(apikey string) {
	ombi.apikey = apikey
}

func (ombi *Ombi) GetUrl() string {
	return ombi.url
}

func (ombi *Ombi) GetKey() string {
	return ombi.apikey
}
