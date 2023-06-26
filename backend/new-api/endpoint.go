package api

import (
	"api/functions"
	"api/links"
	"api/places"
	"api/types"
	"api/user"
	"api/users"
)

// path: /api/*

func HandleEndpoint(endpoint string) (any, error) {
	// (Example) start with: /api/user/995/name
	reducedEndpoint := functions.ReducedFirstLayerEndpoint(endpoint) // (Example) we get: user/995/name/something
	currentEndpoint := functions.CurrentLayer(reducedEndpoint)       // (Example) we get: user

	switch currentEndpoint {
	case "links":
		return links.Get(), nil
	case "user":
		return user.HandleEndpoint(reducedEndpoint)
	case "users":
		return users.Get(), nil
	case "places":
		return places.Get(), nil
	}
	return types.JSONMessageResponse{Success: false, Message: "Invalid endpoint"}, nil
}
