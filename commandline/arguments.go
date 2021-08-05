package commandline

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
)

type Arguments struct {
	ZipCode string
}

func (a *Arguments) Parse() {
	zipCode := flag.String("zip", "", "`zip code` for forecast")
	flag.Parse()

	a.ZipCode = *zipCode
}

func ParseZipCode(maybeZipCode string) (string, error) {
	if maybeZipCode == "" {
		return "", errors.New("please input a zip code")
	}

	match, _ := regexp.Match(`^\d{5}$`, []byte(maybeZipCode))

	if !match {
		return "", fmt.Errorf("invalid zip code: %s", maybeZipCode)
	}

	return maybeZipCode, nil
}
