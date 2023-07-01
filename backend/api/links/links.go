package links

import (
	"errors"
)

func Get() (any, error) {
	if true {
		myData := []string{"https://facebook.dk", "https://google.dk", "https://beksoft.dk"}
		return myData, nil
	}

	return nil, errors.New("database could not connect")
}
