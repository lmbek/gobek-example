package links

import "api/types"

func Get() any {
	myData := []string{"https://facebook.dk", "https://google.dk", "https://beksoft.dk"}

	return types.JSONDataResponse{Success: true, Data: myData}
}
