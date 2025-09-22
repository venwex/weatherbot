package openweather

import "math"

func Convert(k float64) int { // convert from Kelvin to Celcius
	return int(math.Round(k - 273.15))
}