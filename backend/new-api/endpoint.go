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
	// removing "/api/" prefix with endpoint[5:] and then looking for endpoint
	currentEndpoint := functions.NextLayer(endpoint[5:])
	switch currentEndpoint {
	case "links":
		return links.Get(), nil
	case "user":
		return user.HandleEndpoint(currentEndpoint)
	case "users":
		return users.Get(), nil
	case "places":
		return places.Get(), nil
	}
	return types.JSONMessageResponse{Success: false, Message: "Invalid endpoint"}, nil
}
