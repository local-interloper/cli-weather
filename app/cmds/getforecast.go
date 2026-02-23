package cmds

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/fetchers"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
)

func GetForecast(city types.City) tea.Cmd {
	return func() tea.Msg {
		response, err := fetchers.GetForecast(city)

		return msgs.ForecastMsg{
			ForecastResponse: response,
			Err:              err,
		}
	}
}
