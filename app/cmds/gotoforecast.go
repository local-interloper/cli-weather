package cmds

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
)

func GoToForecast(city types.City) tea.Cmd {
	return func() tea.Msg {
		return msgs.GoToForecastMsg{
			City: city,
		}
	}
}
