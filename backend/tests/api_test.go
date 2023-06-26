package tests

import (
	"api"
	"api/types"
	"fmt"
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

		expected := types.JSONMessageResponse{false, "Need identifier: /api/user/{id}"}
		userData, _ := api.HandleEndpoint("/api/user")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call2 /api/user/", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "user id does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call3 /api/user/995", func(test *testing.T) {
		expected := types.JSONDataResponse{true, "995"}
		userData, _ := api.HandleEndpoint("/api/user/995")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call4 /api/user/995/", func(test *testing.T) {
		expected := types.JSONDataResponse{true, "995"}
		userData, _ := api.HandleEndpoint("/api/user/995/")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call5 /api/user/996", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "user id does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/996")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call6 /api/user/996/", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "user id does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/996/")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call7 /api/user/995/name", func(test *testing.T) {
		expected := types.JSONDataResponse{true, "LMB"}
		userData, _ := api.HandleEndpoint("/api/user/995/name")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call8 /api/user/995/name/", func(test *testing.T) {
		expected := types.JSONDataResponse{true, "LMB"}
		userData, _ := api.HandleEndpoint("/api/user/995/name/")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call9 /api/user/996/name", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "user id does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/996/name")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call10 /api/user/996/name/", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "user id does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/996/name/")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call11 /api/user/995/something-something", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "endpoint does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/995/something-something")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
	test.Run("call12 /api/user/996/something-something", func(test *testing.T) {
		expected := types.JSONMessageResponse{false, "endpoint does not exist"}
		userData, _ := api.HandleEndpoint("/api/user/996/something-something")
		result := userData
		//result := userData.(api.JSONDataResponse).Data
		if result != expected {
			test.Errorf("\t\tExpected %v, but got: %v \n", expected, result)
		} else {
			fmt.Printf("\t\tGot expected result: %v \n", expected)
		}
	})
}
