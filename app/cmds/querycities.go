package cmds

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/fetchers"
	"github.com/local-interloper/cli-weather/app/msgs"
)

func QueryCities(query string) tea.Cmd {
	return func() tea.Msg {
		response, err := fetchers.GetCities(query)

		return msgs.CitiesQueryResultMsg{
			Cities: response.Results,
			Ok:     err == nil,
		}
	}
}
