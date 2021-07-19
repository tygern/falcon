package main

import (
	"fmt"
	"github.com/tygern/falcon/commandline"
	"github.com/tygern/falcon/openweathermap"
	"os"
)

func main() {
	fmt.Println("")
	defer fmt.Println("")

	zipCode, err := commandline.ExtractZipCode(os.Args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	forecast := openweathermap.FetchForecast(zipCode)
	commandline.PrintForecast(forecast)
}
