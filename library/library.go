package library

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

func GetLibrarySections(s server.Server) (string, error) {
	url := s.GetUrl() + "/library/sections/"
	method := "GET"

	body, err := request(url, method, s.GetToken())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}

func GetLibraryItems(s server.Server, section string) (string, error) {
	url := s.GetUrl() + "/library/sections/" + section + "/all"
	method := "GET"

	body, err := request(url, method, s.GetToken())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}

func GetLibraryItemMetadata(s server.Server, rating_key string) (string, error) {
	url := s.GetUrl() + "/library/metadata/" + rating_key
	method := "GET"

	body, err := request(url, method, s.GetToken())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}
