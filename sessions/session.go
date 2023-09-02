package sessions

import (
	"net/http"

	xj "github.com/basgys/goxml2json"

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

	body, err := xj.Convert(res.Body)
	if err != nil {
		return "", err
	}

	return body.String(), nil
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

func GetLibraries(s server.Server) (string, error) {
	url := s.GetUrl() + "/library/sections/1/all"
	method := "GET"

	body, err := request(url, method, s.GetToken())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}

func GetAccounts(s server.Server) (string, error) {
	url := s.GetUrl() + "/accounts"
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

func GetTranscodeSessions(s server.Server) (string, error) {
	url := s.GetUrl() + "/transcode/sessions"
	method := "GET"

	body, err := request(url, method, s.GetToken())
	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}

func StopTranscodeSession(s server.Server, sessionKey string) (string, error) {
	url := s.GetUrl() + "/transcode/sessions/" + sessionKey
	method := "DELETE"

	body, err := request(url, method, s.GetToken())
	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}