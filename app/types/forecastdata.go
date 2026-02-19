package types

type ForecastData struct {
	Time                        []string  `json:"time"`
	TemperatureMax              []float32 `json:"temperature_2m_max"`
	TemperatureMin              []float32 `json:"temperature_2m_min"`
	ApparentTemperatureMax      []float32 `json:"apparent_temperature_min"`
	ApparentTemperatureMin      []float32 `json:"apparent_temperature_max"`
	PrecipitationProbabilityMax []float32 `json:"precipitation_probability_max"`
	WeatherCode                 []int     `json:"weather_code"`
}
