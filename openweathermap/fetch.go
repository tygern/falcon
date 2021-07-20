package openweathermap

import (
	"encoding/json"
	"fmt"
	"github.com/tygern/falcon/forecast"
	"io/ioutil"
	"net/http"
)

type OWMClient struct {
	ApiKey     string
	BaseUrl    string
	RestClient interface {
		Get(string) (*http.Response, error)
	}
}

func (client OWMClient) FetchForecast(zipCode string) (*forecast.Forecast, error) {
	response, err := client.callForecast(zipCode)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data owmResponse
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, err
	}

	return &forecast.Forecast{ZipCode: zipCode, CurrentTemperature: kelvinToFahrenheit(data.Main.Temp)}, nil
}

func (client OWMClient) callForecast(zipCode string) (*http.Response, error) {
	url := fmt.Sprintf("%s/data/2.5/weather?zip=%s,us&appid=%s", client.BaseUrl, zipCode, client.ApiKey)
	response, err := client.RestClient.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("error when fetching forecast for %s", zipCode)
	}
	return response, nil
}

func kelvinToFahrenheit(kelvin float32) float32 {
	return 1.8*(kelvin-273) + 32
}

type owmResponse struct {
	Main owmMain
}

type owmMain struct {
	Temp float32
}
