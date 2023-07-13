package api

import (
	"api/endpoints/example"
	"api/endpoints/links"
	"api/endpoints/native"
	"api/endpoints/places"
	"api/endpoints/profile"
	"api/endpoints/profiles"
	"api/functions"
	"errors"
	"net/http"
)

// Example data (path): /api/user/995/name/something
// HandleEndpoint - This method can be mimiced, it reduces the endpoint and selects the current for use
func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	reducedEndpoint := functions.ReducedFirstLayerEndpoint(endpoint) // Example data: user/995/name/something
	currentEndpoint := functions.CurrentLayer(reducedEndpoint)       // Example data: user

	switch currentEndpoint {
	case "example":
		return example.HandleEndpoint(reducedEndpoint, response, request)
	case "links":
		return links.HandleEndpoint(reducedEndpoint, response, request)
	case "native":
		return native.HandleEndpoint(reducedEndpoint, response, request)
	case "places":
		return places.HandleEndpoint(reducedEndpoint, response, request)
	case "profile":
		return profile.HandleEndpoint(reducedEndpoint, response, request)
	case "profiles":
		return profiles.HandleEndpoint(reducedEndpoint, response, request)
	}

	return nil, errors.New("invalid endpoint")
}
