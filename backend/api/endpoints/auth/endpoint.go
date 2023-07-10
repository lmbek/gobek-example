package auth

import (
	"api/endpoints/auth/session"
	"api/endpoints/auth/signin"
	"api/endpoints/auth/signout"
	"api/functions"
	"errors"
	"net/http"
)

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	_, current := functions.StandardEndpointData(endpoint)

	switch current {
	case "active":
		return session.HandleEndpoint(current, response, request)
	case "signin":
		return signin.HandleEndpoint(current, response, request)
	case "signout":
		return signout.HandleEndpoint(current, response, request)
	}

	return nil, errors.New("You can use the following: /api/auth/signin, /api/auth/signout/, /api/auth/signup")
}
