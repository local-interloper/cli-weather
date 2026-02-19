package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/local-interloper/cli-weather/app/msgs"
)

func RowsFromForecastMsg(msg msgs.ForecastMsg) []table.Row {
	rows := []table.Row{}

	for i := range len(msg.Daily.Time) {
		row := []string{}

		row = append(row, fmt.Sprintf("%s", IsoToWeekday(msg.Daily.Time[i])))
		row = append(row, fmt.Sprintf("%s", CodeToStatus(msg.Daily.WeatherCode[i])))
		row = append(row, fmt.Sprintf("%.1f", msg.Daily.PrecipitationProbabilityMax[i]))
		row = append(row, fmt.Sprintf("%.1f", msg.Daily.TemperatureMin[i]))
		row = append(row, fmt.Sprintf("%.1f", msg.Daily.TemperatureMax[i]))

		rows = append(rows, row)
	}

	return rows
}
