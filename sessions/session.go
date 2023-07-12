package sessions

import (
	"io/ioutil"
	"net/http"

	"github.com/aashish2057/plexgoutils/server"
)

func request(url string, method string, token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("X-Plex-Token", token)

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

func GetActiveSessions(s server.Server) (string, error) {
	url := s.GetUrl() + "/status/sessions"
	method := "GET"

	body, err := request(url, method, s.GetToken())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}

func GetSessionHistory(s server.Server) (string, error) {
	url := s.GetUrl() + "/status/sessions/history/all"
	method := "GET"

	body, err := request(url, method, s.GetToken())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}
