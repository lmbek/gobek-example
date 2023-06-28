package id

import (
	"api/functions"
	"api/types"
	"api/user/id/age"
	"api/user/id/name"
)

func HandleIdentifierEndpoint(endpoint string) (any, error) {
	reducedEndpoint := functions.ReducedEndpoint(endpoint)
	identifier := functions.CurrentLayer(reducedEndpoint) // find identifier and use as UserID

	if len(identifier) >= 1 {
		userID, err := Get(identifier)

		if err != nil {
			// user do not exist or database could not connect, etc...
			return types.JSONMessageResponse{Success: false, Message: err.Error()}, nil
		}

		// check if there is more after identifier - if there is more endpoints
		reducedEndpoint = functions.ReducedEndpoint(reducedEndpoint) // (Example) we get: name/something
		currentEndpoint := functions.CurrentLayer(reducedEndpoint)   // (Example) we get: name

		// if we could find user but there was not requested any endpoints after identifier: (example) 995
		if len(currentEndpoint) == 0 {
			return types.JSONDataResponse{Success: true, Data: identifier}, nil
		}

		switch currentEndpoint {
		// path: {id}/name/something
		case "name":
			return name.Get(userID), nil
		case "age":
			return age.Get(userID), nil
		}

		return types.JSONMessageResponse{Success: false, Message: "Invalid endpoint"}, nil
	}
	return types.JSONMessageResponse{Success: false, Message: "Need identifier: /api/user/{id}"}, nil
}
