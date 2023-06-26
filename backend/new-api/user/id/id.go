package id

import "api/types"

func Get(identifier string) any {
	if identifier == "995" {
		return types.JSONDataResponse{true, "995"}
	} else {
		return types.JSONMessageResponse{false, "user id does not exist"}
	}
}
