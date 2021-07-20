package main

import (
	"fmt"
	"github.com/tygern/falcon/commandline"
	"github.com/tygern/falcon/openweathermap"
	"github.com/tygern/falcon/restsupport"
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

	client := openweathermap.OWMClient{
		ApiKey:     os.Getenv("OWM_KEY"),
		BaseUrl:    "https://api.openweathermap.org",
		RestClient: restsupport.RestTemplate{},
	}

	result, err := client.FetchForecast(zipCode)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	commandline.PrintForecast(*result)
}
