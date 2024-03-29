package tests

import (
	"api"
	"api/functions"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIIllegalPath(test *testing.T) {
	test.Run("1_something/something", func(test *testing.T) {
		// endpoint
		endpoint := "something/something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := "invalid endpoint"
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
func TestAPIUsersEndpoint(test *testing.T) {
	test.Run("1_/api/profiles", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profiles"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := []string{"995"}
		result, _ := api.HandleEndpoint(endpoint, response, request)
		if !reflect.DeepEqual(result, expected) {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("2_/api/profiles/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profiles/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := []string{"995"}
		result, _ := api.HandleEndpoint(endpoint, response, request)
		fmt.Println(result)
		if !reflect.DeepEqual(result, expected) {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("3_/api/profiles/something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profiles/something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)
		// EXPECTED
		expected := errors.New("invalid endpoint").Error()

		// RESULT
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("4_/api/profiles/something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profiles/something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)
		// EXPECTED
		expected := errors.New("invalid endpoint").Error()

		// RESULT
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
func TestAPILinksEndpoint(test *testing.T) {
	test.Run("1_/api/links", func(test *testing.T) {
		// endpoint
		endpoint := "/api/links"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := []string{"https://google.dk", "https://facebook.com", "https://linkedin.com"}

		// RESULT
		result, _ := api.HandleEndpoint(endpoint, response, request)

		// GIVE TEST RESULT
		if !reflect.DeepEqual(result, expected) {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("2_/api/links/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/links/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := []string{"https://google.dk", "https://facebook.com", "https://linkedin.com"}

		// RESULT
		result, _ := api.HandleEndpoint(endpoint, response, request)

		// GIVE TEST RESULT
		if !reflect.DeepEqual(result, expected) {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
func TestAPIPlacesEndpoint(test *testing.T) {
	test.Run("1_/api/places", func(test *testing.T) {
		// endpoint
		endpoint := "/api/places"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := []string{"København", "Odense", "Aarhus"}

		// RESULT
		result, _ := api.HandleEndpoint(endpoint, response, request)

		// GIVE TEST RESULT
		if !reflect.DeepEqual(result, expected) {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("2_/api/places/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/places/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		// EXPECTED
		expected := []string{"København", "Odense", "Aarhus"}

		// RESULT
		result, _ := api.HandleEndpoint(endpoint, response, request)

		// GIVE TEST RESULT
		if !reflect.DeepEqual(result, expected) {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
func TestAPIUserEndpoint(test *testing.T) {
	test.Run("1_/api/profile", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("need identifier: /api/profile/{id}").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("2_/api/profile/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("need identifier: /api/profile/{id}").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("3_/api/profile/995", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := "995"
		result, _ := api.HandleEndpoint(endpoint, response, request)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("4_/api/profile/995/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := "995"
		result, _ := api.HandleEndpoint(endpoint, response, request)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("5_/api/profile/996", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("6_/api/profile/996/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("7_/api/profile/995/name", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/name"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := "LMB"
		result, _ := api.HandleEndpoint(endpoint, response, request)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("8_/api/profile/995/name/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/name/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := "LMB"
		result, _ := api.HandleEndpoint(endpoint, response, request)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("9_/api/profile/996/name", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/name"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("10_/api/profile/996/name/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/name/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("11_/api/profile/995/name/something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/name/something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("12_/api/profile/995/name/something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/name/something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("13_/api/profile/996/name/something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/name/something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("14_/api/profile/996/name/something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/name/something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("15_/api/profile/995/something-something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/something-something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("16_/api/profile/995/something-something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/something-something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint(endpoint, response, request)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("17_/api/profile/996/something-something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/something-something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		_, err := api.HandleEndpoint(endpoint, response, request)

		// EXPECTED
		expected := functions.InvalidEndpointError().Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("18_/api/profile/996/something-something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/something-something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", endpoint, nil)

		_, err := api.HandleEndpoint(endpoint, response, request)

		// EXPECTED
		expected := functions.InvalidEndpointError().Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("19_/api/profile/995/something-something/more-something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/something-something/more-something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", endpoint, nil)

		_, err = api.HandleEndpoint(endpoint, response, request)

		// EXPECTED
		expected := functions.InvalidEndpointError().Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("20_/api/profile/995/something-something/more-something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/995/something-something/more-something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", endpoint, nil)

		_, err = api.HandleEndpoint(endpoint, response, request)

		// EXPECTED
		expected := functions.InvalidEndpointError().Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("21_/api/profile/996/something-something/more-something", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/something-something/more-something"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", endpoint, nil)

		_, err = api.HandleEndpoint(endpoint, response, request)

		// EXPECTED
		expected := functions.InvalidEndpointError().Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("22_/api/profile/996/something-something/more-something/", func(test *testing.T) {
		// endpoint
		endpoint := "/api/profile/996/something-something/more-something/"

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", endpoint, nil)

		_, err = api.HandleEndpoint(endpoint, response, request)

		// EXPECTED
		expected := functions.InvalidEndpointError().Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
