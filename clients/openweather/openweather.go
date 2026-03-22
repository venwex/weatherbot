package openweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenWeatherClient struct {
	apiKey string
}

func New(apiKey string) *OpenWeatherClient {
	return &OpenWeatherClient{
		apiKey: apiKey,
	}
}

func (o OpenWeatherClient) Coordinates(city string) (Coordinates, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s", city, o.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error getting coordinates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // 200 OK - successfull response
		return Coordinates{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var coordinatesResponse []CoordinatesResponse

	err = json.NewDecoder(resp.Body).Decode(&coordinatesResponse)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error parsing response body: %w", err)
	}

	if len(coordinatesResponse) == 0 {
		return Coordinates{}, fmt.Errorf("error empty coordinates: %w", err)
	}

	return Coordinates {
		Lat: coordinatesResponse[0].Lat,
		Lon: coordinatesResponse[0].Lon,
	}, nil
}

func (o OpenWeatherClient) Weather(c Coordinates) (Weather, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s",
    c.Lat, c.Lon, o.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, fmt.Errorf("error requesting weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var weatherResponse WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return Weather{}, fmt.Errorf("error parsing response body: %w", err)
	}

	return Weather{
		Temp: weatherResponse.Main.Temp,
		Description: weatherResponse.Weather[0].Description,
	}, nil
}