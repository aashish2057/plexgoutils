package sessions

import (
	"io/ioutil"
	"net/http"

	"github.com/aashish2057/plexgoutils/server"
)

func GetActiveSessions(s server.Server) (string, error) {
	url := s.GetUrl() + "/status/sessions"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("X-Plex-Token", s.GetToken())

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
