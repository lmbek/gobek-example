package functions

import "strings"

func CurrentLayer(endpoint string) string {
	// /api/personal//user/995/name

	splitted := strings.Split(endpoint, "/")
	// ["", "api", "user", "995", "name"]
	if len(splitted) >= 1 {
		return splitted[0] // /personal//user/995/name
	}
	return ""
}

func ReducedEndpoint(endpoint string) string {
	splitted := strings.Split(endpoint, "/")
	if len(splitted) >= 1 {
		return strings.Join(splitted[1:], "/")
	}
	return ""
}

func ReducedFirstLayerEndpoint(endpoint string) string {
	splitted := strings.Split(endpoint, "/")
	if len(splitted) >= 2 {
		return strings.Join(splitted[2:], "/")
	}
	return ""
}
