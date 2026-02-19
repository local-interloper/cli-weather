package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/local-interloper/cli-weather/app/msgs"
)

func ColumnsFromForecastMsg(msg msgs.ForecastMsg) []table.Column {
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
			Title: fmt.Sprintf("Rain %s", msg.DailyUnits.PrecipitationProbabilityMax),
			Width: 8,
		},
		{
			Title: fmt.Sprintf("Min %s", msg.DailyUnits.TemperatureMin),
			Width: 8,
		},
		{
			Title: fmt.Sprintf("Max %s", msg.DailyUnits.TemperatureMax),
			Width: 8,
		},
	}
}
