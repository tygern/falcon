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

	arguments := new(commandline.Arguments)
	arguments.Parse()

	validZipCode, err := commandline.ParseZipCode(arguments.ZipCode)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := openweathermap.OWMClient{
		ApiKey:     os.Getenv("OWM_KEY"),
		BaseUrl:    "https://api.openweathermap.org",
		RestClient: restsupport.RestTemplate{},
	}

	result, err := client.FetchForecast(validZipCode)

	if err != nil {
		fmt.Println(err)
		return
	}

	commandline.PrintForecast(*result)
}
