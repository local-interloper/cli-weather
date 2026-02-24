package utils

import "fmt"

var codemap map[int]string = map[int]string{
	// Clear
	0: "☀️ Clear",
	1: "🌤️ Mainly clear",
	2: "⛅ Partly cloudy",
	3: "🌥️ Overcast",

	// Fog
	45: "🌫️ Fog",
	48: "🌫️ Depositing rime fog",

	// Drizzle
	51: "🌦️ Light drizzle",
	52: "🌦️ Moderate drizzle",
	53: "🌦️ Dense drizzle",
	56: "🌦️️ Light freezing drizzle",
	57: "️🌦️ Heavy freezing drizzle",

	// Rain
	61: "🌧️ Slight rain",
	63: "🌧️ Moderate rain",
	65: "🌧️ Heavy rain",
	66: "🌧️ Light freezing rain",
	67: "🌧️ Heavy freezing rain",

	// Snow
	71: "🌨️ Slight snowfall",
	73: "🌨️ Moderate snowfall",
	75: "🌨️ Heavy snowfall",
	77: "🌨️ Snow grains",

	// Rain showers
	80: "🌧️ Slight rain shower",
	81: "🌧️ Moderate rain shower",
	82: "🌧️ Violent rain shower",

	// Snow showers
	85: "🌨️ Slight snow shower",
	86: "🌨️ Heavy snow shower",

	// Thunderstorm
	95: "⛈️  Thunderstorm",
	96: "⛈️  Thunderstorm w/ slight hail",
	99: "⛈️  Thunderstorm w/ heavy hail",
}

func CodeToStatus(code int) string {
	status, ok := codemap[code]

	if !ok {
		return fmt.Sprintf("Code %d not implemented", code)
	}

	return status
}
