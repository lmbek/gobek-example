package api

import (
	"api/example"
	"api/functions"
	"api/links"
	"api/native"
	"api/places"
	"api/user"
	"api/users"
	"errors"
)

// path: /api/*

func HandleEndpoint(endpoint string) (any, error) {
	// (Example) start with: /api/user/995/name
	reducedEndpoint := functions.ReducedFirstLayerEndpoint(endpoint) // (Example) we get: user/995/name/something
	currentEndpoint := functions.CurrentLayer(reducedEndpoint)       // (Example) we get: user

	switch currentEndpoint {
	case "native":
		return native.HandleEndpoint(reducedEndpoint)
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
