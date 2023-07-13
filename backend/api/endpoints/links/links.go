package links

import (
	"errors"
)

func Get() (any, error) {
	// if we have data, send data
	if true {
		myData := []string{"https://google.dk", "https://facebook.com", "https://linkedin.com"}
		return myData, nil
	}

	// else send the error (modify current error)
	return nil, errors.New("database could not connect")
}
