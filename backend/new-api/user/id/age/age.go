package age

import "api/types"

func Get(id string) any {
	if id == "995" {
		return types.JSONDataResponse{true, 28}
	} else {
		return types.JSONMessageResponse{false, "user id does not exist"}
	}
}
