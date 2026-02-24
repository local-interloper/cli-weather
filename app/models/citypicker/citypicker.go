package citypicker

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/cmds"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
	"github.com/local-interloper/cli-weather/app/utils"
)

type Model struct {
	query      string
	citiesList list.Model
	cities     []types.City
	help       help.Model
}

func New(query string) Model {
	width, height := utils.TermSize()

	ls := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height-1)
	ls.Title = "Select a city"
	ls.InfiniteScrolling = true
	ls.SetShowHelp(false)

	return Model{
		query:      query,
		citiesList: ls,
		help:       help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(cmds.QueryCities(m.query))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.CitiesQueryResultMsg:
		m.cities = msg.Cities
		items := utils.CitiesToItems(m.cities)
		m.citiesList.SetItems(items)
		m.citiesList.FilterInput.Focus()
		return m, nil

	case tea.KeyMsg:
		if key.Matches(msg, keymap.Accept) {
			if len(m.cities) == 0 {
				return m, nil
			}

			index := m.citiesList.Index()
			city := m.cities[index]
			utils.SetDefaultCity(city)
			return m, cmds.GoToForecast(city)
		}

		if key.Matches(msg, keymap.Back) {
			return m, cmds.GoToCityQuery()
		}
	}

	{
		var cmd tea.Cmd
		if m.citiesList, cmd = m.citiesList.Update(msg); cmd != nil {
			return m, cmd
		}
	}

	return m, nil
}

func (m Model) View() string {
	out := ""
	out += m.citiesList.View()
	out += "\n"
	out += m.help.View(keymap)
	return out
}
