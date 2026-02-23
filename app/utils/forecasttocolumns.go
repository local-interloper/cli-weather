package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/local-interloper/cli-weather/app/types"
)

func ForecastToColumns(forecast types.ForecastResponse) []table.Column {
	return []table.Column{
		{
			Title: "Date",
			Width: 10,
		},
		{
			Title: "Status",
			Width: 25,
		},
		{
			Title: fmt.Sprintf("Rain %s", forecast.DailyUnits.PrecipitationProbabilityMax),
			Width: 8,
		},
		{
			Title: fmt.Sprintf("Min %s", forecast.DailyUnits.TemperatureMin),
			Width: 8,
		},
		{
			Title: fmt.Sprintf("Max %s", forecast.DailyUnits.TemperatureMax),
			Width: 8,
		},
	}
}
