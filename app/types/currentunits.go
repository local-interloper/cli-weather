package types

type CurrentUnits struct {
	Time                string `json:"time"`
	Interval            string `json:"interval"`
	Temperature         string `json:"temperature_2m"`
	ApparentTemperature string `json:"apparent_temperature"`
}
