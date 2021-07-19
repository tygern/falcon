package commandline

import (
	"fmt"
	"github.com/tygern/falcon/forecast"
)

func PrintForecast(forecast forecast.Forecast) {
	fmt.Printf("Weather for %v\n", forecast.ZipCode)
	fmt.Printf("Current temperature: %d degrees\n", forecast.CurrentTemperature)
}
