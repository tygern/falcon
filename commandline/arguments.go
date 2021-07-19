package commandline

import (
	"errors"
	"fmt"
	"regexp"
)

func ExtractZipCode(args []string) (zipCode string, err error) {
	if len(args) < 2 {
		err = errors.New("please input a zip code")
		return
	}

	zipCode = args[1]

	match, _ := regexp.Match(`^\d{5}$`, []byte(zipCode))

	if !match {
		err = fmt.Errorf("invalid zip code: %v", zipCode)
		return
	}

	return
}
