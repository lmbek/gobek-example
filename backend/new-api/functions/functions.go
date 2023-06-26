package functions

import "strings"

func NextLayer(endpoint string) string {
	splitted := strings.Split(endpoint, "/")
	if len(splitted) > 0 {
		return splitted[0] // identifier reached, maybe?
	}
	return ""
}

func CuttedEndpoint(endpoint string) string {
	splitted := strings.Split(endpoint, "/")
	if len(splitted) > 0 {
		return strings.Join(splitted[1:], "/")
	}
	return ""
}
