package cmds

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/fetchers"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
)

func GetForecast(cityName string) tea.Cmd {
	return func() tea.Msg {
		city := fetchers.GetCity(cityName)

		if city == nil {
			return msgs.ForecastMsg{
				Err: errors.New("Failed to find city"),
			}
		}

		coords := types.Coords{
			Latitude:  city.Latitude,
			Longitude: city.Longitude,
		}

		return fetchers.GetForecast(coords)
	}
}
