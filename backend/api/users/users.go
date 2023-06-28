package users

import "api/types"

func Get() any {
	// receive data from database
	myData := []string{"Lars", "ida", "emil"}

	return types.JSONDataResponse{true, myData}
}
