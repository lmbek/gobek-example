package native

import (
	"api/functions"
	"api/native/choose"
	"errors"
)

func HandleEndpoint(endpoint string) (any, error) {
	reducedEndpoint := functions.ReducedEndpoint(endpoint)
	currentEndpoint := functions.CurrentLayer(reducedEndpoint)

	// This is overwriting current endpoint, such as if we want some other features extra...
	switch currentEndpoint {
	case "folder":
		return choose.Folder()
	}

	return nil, errors.New("invalid endpoint")
}
