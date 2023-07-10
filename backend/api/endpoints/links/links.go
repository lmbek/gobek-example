package links

import (
	"errors"
)

func Get() (any, error) {
	if true {
		myData := []string{"https://google.dk", "https://facebook.com", "https://linkedin.com"}
		return myData, nil
	}

	return nil, errors.New("database could not connect")
}
