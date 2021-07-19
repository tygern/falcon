package openweathermap

import (
	"github.com/tygern/falcon/forecast"
)

func FetchForecast(zipCode string) forecast.Forecast {
	return forecast.Forecast{ZipCode: zipCode}
}
