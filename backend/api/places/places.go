package places

import "api/types"

func Get() any {
	myData := []string{"Grenaa", "Odense", "Roskilde"}

	return types.JSONDataResponse{Success: true, Data: myData}
}
