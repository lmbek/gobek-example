package user

import (
	"api/functions"
	"api/types"
	"api/user/age"
	"api/user/name"
	"strings"
)

// HandleEndpoint - If no endpoints, then we try to return all data on user (if user id exists),
// if they dont have a session, some or all parts of the api might not be returned
// if no user id with that identifier, then an error with "no such user" will be given
func HandleEndpoint(endpoint string) (any, error) {
	//fmt.Printf("Endpoint: %v \n", endpoint)
	currentEndpoint := functions.NextLayer(endpoint)
	//fmt.Printf("Current: %v \n", currentEndpoint)
	identifier := functions.NextLayer(currentEndpoint)
	cutted := functions.CuttedEndpoint(identifier) // */...
	//fmt.Printf("Cutted: %v \n", cutted)
	// set current endpoint
	currentEndpoint = functions.NextLayer(cutted)
	// path: /api/user/{id}
	// Identifier: {id}
	if strings.Contains(endpoint, "/") == false {
		// give error: we need identifier
		return types.JSONMessageResponse{Success: false, Message: "Need identifier: /api/user/{id}"}, nil
	}
	// we got identifier, store it in userId variable
	userId := currentEndpoint
	//cutted = functions.CuttedEndpoint(cutted)
	currentEndpoint = functions.NextLayer(cutted)

	switch currentEndpoint {
	// path: /api/user/{id}/name
	case "name":
		return name.Get(userId), nil
	case "age":
		return age.Get(userId), nil
	}

	return types.JSONMessageResponse{Success: false, Message: "Invalid endpoint"}, nil
}
