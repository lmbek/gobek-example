package native

import (
	"api/endpoints/native/folder"
	"api/functions"
	"net/http"
)

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	err := functions.HandleAsLastEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	reducedEndpoint, currentEndpoint := functions.StandardEndpointData(endpoint)

	switch currentEndpoint {
	case "folder":
		return folder.HandleEndpoint(reducedEndpoint, response, request)
	}

	return functions.InvalidEndpoint()
}
