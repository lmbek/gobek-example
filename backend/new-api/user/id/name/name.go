package name

import "api/types"

func Get(id string) any {
	if id == "995" {
		return types.JSONDataResponse{true, "LMB"}
	} else {
		return types.JSONMessageResponse{false, "user id does not exist"}
	}
}
