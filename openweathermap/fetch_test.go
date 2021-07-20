package openweathermap

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tygern/falcon/forecast"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestFetchForecast(t *testing.T) {
	mockClient := MockRestClient{
		Body:       "{\"main\": {\"temp\": 300, \"other\": \"to ignore\"}}",
		StatusCode: 200,
	}

	client := OWMClient{
		ApiKey:     "super-secret",
		BaseUrl:    "https://owm.example.com",
		RestClient: &mockClient,
	}

	result, err := client.FetchForecast("12345")

	assert.Equal(t, &forecast.Forecast{ZipCode: "12345", CurrentTemperature: 80.6}, result)
	assert.Nil(t, err)
	assert.Equal(t, "https://owm.example.com/data/2.5/weather?zip=12345,us&appid=super-secret", mockClient.RequestUrl)
}

func TestFetchForecast400(t *testing.T) {
	mockClient := MockRestClient{
		Body:       "not json",
		StatusCode: 400,
	}

	client := OWMClient{
		ApiKey:     "super-secret",
		BaseUrl:    "https://owm.example.com",
		RestClient: &mockClient,
	}

	_, err := client.FetchForecast("12345")

	assert.Equal(t, errors.New("error when fetching forecast for 12345"), err)
}

func TestFetchForecastError(t *testing.T) {
	mockClient := MockRestClient{Error: errors.New("uh oh")}
	client := OWMClient{RestClient: &mockClient}

	_, err := client.FetchForecast("12345")

	assert.Equal(t, errors.New("uh oh"), err)
}

func TestFetchForecastJsonError(t *testing.T) {
	mockClient := MockRestClient{
		Body:       "not json",
		StatusCode: 200,
	}

	client := OWMClient{
		ApiKey:     "super-secret",
		BaseUrl:    "https://owm.example.com",
		RestClient: &mockClient,
	}
	_, err := client.FetchForecast("12345")

	assert.IsType(t, &json.SyntaxError{}, err)
}

type MockRestClient struct {
	Body       string
	StatusCode int
	Error      error
	RequestUrl string
}

func (c *MockRestClient) Get(url string) (*http.Response, error) {
	c.RequestUrl = url

	if c.Error != nil {
		return nil, c.Error
	}

	response := http.Response{
		StatusCode: c.StatusCode,
		Body:       ioutil.NopCloser(strings.NewReader(c.Body)),
	}

	return &response, nil
}
