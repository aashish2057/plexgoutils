package ombi_server

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

	req.Header.Add("ApiKey", token)

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

func GetAllMovieRequests(o server.Ombi) (string, error) {
	url := o.GetUrl() + "/api/v1/Request/movie"
	method := "GET"

	body, err := request(url, method, o.GetKey())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}

func GetAllTvRequests(o server.Ombi) (string, error) {
	url := o.GetUrl() + "/api/v1/Request/tv"
	method := "GET"

	body, err := request(url, method, o.GetKey())

	if err != nil {
		return "", err
	} else {
		return body, nil
	}
}
