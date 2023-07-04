package id

import (
	"api/functions"
	"api/user/id/age"
	"api/user/id/name"
	"errors"
)

func HandleIdentifierEndpoint(endpoint string) (any, error) {
	reducedEndpoint := functions.ReducedEndpoint(endpoint)
	identifier := functions.CurrentLayer(reducedEndpoint) // find identifier and use as UserID

	if len(identifier) >= 1 {
		userID, err := Get(identifier)

		if err != nil {
			// user do not exist or database could not connect, etc...
			return nil, err
		}

		// check if there is more after identifier - if there is more endpoints
		reducedEndpoint = functions.ReducedEndpoint(reducedEndpoint) // (Example) we get: name/something
		currentEndpoint := functions.CurrentLayer(reducedEndpoint)   // (Example) we get: name

		// if we could find user but there was not requested any endpoints after identifier: (example) 995
		if len(currentEndpoint) == 0 {
			return identifier, nil
		}

		switch currentEndpoint {
		// path: {id}/name/something
		case "name":
			return name.Get(userID)
		case "age":
			return age.Get(userID)
		}

		return nil, errors.New("invalid endpoint")
	}

	return nil, errors.New("need identifier: /api/user/{id}")
}
