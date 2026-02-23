package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/local-interloper/cli-weather/app/types"
)

func ForecastToRows(forecast types.ForecastResponse) []table.Row {
	rows := []table.Row{}

	for i := range len(forecast.Daily.Time) {
		row := []string{}

		row = append(row, fmt.Sprintf("%s", IsoToWeekday(forecast.Daily.Time[i])))
		row = append(row, fmt.Sprintf("%s", CodeToStatus(forecast.Daily.WeatherCode[i])))
		row = append(row, fmt.Sprintf("%.1f", forecast.Daily.PrecipitationProbabilityMax[i]))
		row = append(row, fmt.Sprintf("%.1f", forecast.Daily.TemperatureMin[i]))
		row = append(row, fmt.Sprintf("%.1f", forecast.Daily.TemperatureMax[i]))

		rows = append(rows, row)
	}

	return rows
}
