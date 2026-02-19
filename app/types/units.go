package types

type Units struct {
	Time                        string `json:"time"`
	TemperatureMax              string `json:"temperature_2m_max"`
	TemperatureMin              string `json:"temperature_2m_min"`
	ApparentTemperatureMax      string `json:"apparent_temperature_min"`
	ApparentTemperatureMin      string `json:"apparent_temperature_max"`
	PrecipitationProbabilityMax string `json:"precipitation_probability_max"`
	WeatherCode                 string `json:"weather_code"`
}
