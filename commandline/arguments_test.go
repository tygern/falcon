package commandline

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractZipCode(t *testing.T) {
	zipCode, err := ParseZipCode("12345")

	assert.Equal(t, "12345", zipCode)
	assert.Nil(t, err)
}

func TestExtractZipCodeMissing(t *testing.T) {
	_, err := ParseZipCode("")

	assert.Equal(t, errors.New("please input a zip code"), err)
}

func TestExtractZipCodeWrongFormat(t *testing.T) {
	_, err := ParseZipCode("a")

	assert.Equal(t, errors.New("invalid zip code: a"), err)
}
