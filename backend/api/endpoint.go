package api

import (
	"api/download"
	"api/example"
	"api/functions"
	"api/html"
	"api/links"
	"api/native"
	"api/places"
	"api/png"
	"api/user"
	"api/users"
	"errors"
	"net/http"
)

// path: /api/*

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	// (Example) start with: /api/user/995/name
	reducedEndpoint := functions.ReducedFirstLayerEndpoint(endpoint) // (Example) we get: user/995/name/something
	currentEndpoint := functions.CurrentLayer(reducedEndpoint)       // (Example) we get: user

	switch currentEndpoint {
	case "download":
		return download.Start(response, request)
	case "native":
		return native.HandleEndpoint(reducedEndpoint)
	case "png":
		return png.Get(response, request)
	case "html":
		return html.Get(response, request)
	case "example":
		return example.List(true)
	case "links":
		return links.Get()
	case "user":
		return user.HandleEndpoint(reducedEndpoint)
	case "users":
		return users.Get()
	case "places":
		return places.Get()
	}
	return nil, errors.New("invalid endpoint")
}
