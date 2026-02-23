package types

type Current struct {
	Time                string  `json:"time"`
	Interval            int     `json:"interval"`
	Temperature         float32 `json:"temperature_2m"`
	ApparentTemperature float32 `json:"apparent_temperature"`
}
