package openweathermap

import (
	"github.com/stretchr/testify/assert"
	"github.com/tygern/falcon/forecast"
	"testing"
)

func TestFetchForecast(t *testing.T) {
	result := FetchForecast("12345")

	assert.Equal(t, forecast.Forecast{ZipCode: "12345"}, result)
}
