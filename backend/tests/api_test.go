package tests

import (
	"api"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	func TestAPIUsersEndpoint(test *testing.T) {
		test.Run("call1 /api/links", func(test *testing.T) {
			expected := types.JSONMessageResponse{true, "[https://facebook.dk, https://google.dk, https://beksoft.dk]"}.Message
			userData := api.HandleEndpoint("/api/links").(types.JSONDataResponse).Data
			result := userData
			//result := userData.(api.JSONDataResponse).Data
			if result != expected {
				test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
			} else {
				fmt.Printf("\t\tGot expected result: %v \n", expected)
			}
		})
	}
*/
func TestAPIUserEndpoint(test *testing.T) {
	test.Run("call1 /api/user", func(test *testing.T) {

		expected := errors.New("need identifier: /api/user/{id}").Error()
		_, err := api.HandleEndpoint("/api/user", nil, nil)
		result := err.Error()
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call2 /api/user/", func(test *testing.T) {
		expected := errors.New("need identifier: /api/user/{id}").Error()
		_, err := api.HandleEndpoint("/api/user/", nil, nil)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call3 /api/user/995", func(test *testing.T) {
		expected := "995"
		result, _ := api.HandleEndpoint("/api/user/995", nil, nil)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call4 /api/user/995/", func(test *testing.T) {
		expected := "995"
		result, _ := api.HandleEndpoint("/api/user/995/", nil, nil)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call5 /api/user/996", func(test *testing.T) {
		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint("/api/user/996", nil, nil)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call6 /api/user/996/", func(test *testing.T) {
		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint("/api/user/996/", nil, nil)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call7 /api/user/995/name", func(test *testing.T) {
		expected := "LMB"
		result, _ := api.HandleEndpoint("/api/user/995/name", nil, nil)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call8 /api/user/995/name/", func(test *testing.T) {
		expected := "LMB"
		result, _ := api.HandleEndpoint("/api/user/995/name/", nil, nil)

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call9 /api/user/996/name", func(test *testing.T) {
		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint("/api/user/996/name", nil, nil)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call10 /api/user/996/name/", func(test *testing.T) {
		expected := errors.New("user id does not exist").Error()
		_, err := api.HandleEndpoint("/api/user/996/name/", nil, nil)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call11 /api/user/995/something-something", func(test *testing.T) {
		expected := errors.New("invalid endpoint").Error()
		_, err := api.HandleEndpoint("/api/user/995/something-something", nil, nil)
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call12 /api/user/996/something-something", func(test *testing.T) {
		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/html", nil)

		_, err = api.HandleEndpoint("/api/user/996/something-something", response, request)

		// EXPECTED
		expected := errors.New("user id does not exist").Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
func TestAPIHTMLEndpoint(test *testing.T) {
	test.Run("call1 /api/html/", func(test *testing.T) {
		// EXPECTED
		expected := errors.New("not json response").Error()

		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/html/", nil)

		_, err = api.HandleEndpoint("/api/html/", response, request)

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call8 /api/html", func(test *testing.T) {
		// create a mock response writer & request
		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/html", nil)

		_, err = api.HandleEndpoint("/api/html", response, request)

		// EXPECTED
		expected := errors.New("not json response").Error()

		// RESULT
		result := err.Error()

		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
