package signin

import (
	"api/functions"
	"net/http"
)

func HandleEndpoint(endpoint string, response http.ResponseWriter, request *http.Request) (any, error) {
	err := functions.HandleAsLastEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	switch request.Method {
	case http.MethodGet:
		return nil, functions.NotSupportedError()
	case http.MethodPost:
		return SignIn(response, request)
	case http.MethodPut:
		return nil, functions.NotSupportedError()
	case http.MethodDelete:
		return nil, functions.NotSupportedError()
	}

	return functions.InvalidEndpoint()
}
