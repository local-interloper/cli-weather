package forecast

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/local-interloper/cli-weather/app/cmds"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
	"github.com/local-interloper/cli-weather/app/utils"
)

type Model struct {
	city   types.City
	banner string
	table  table.Model
	help   help.Model
}

func New(city types.City) Model {
	width, height := utils.TermSize()

	banner := fmt.Sprintf("Weather forecast for: %s, %s\n\n", city.Name, city.Country)

	ts := table.DefaultStyles()
	ts.Header = ts.Header.Background(lipgloss.Color("#8888ff"))
	ts.Selected = ts.Selected.Foreground(lipgloss.NoColor{}).Bold(false)

	tb := table.New(
		table.WithWidth(width),
		table.WithHeight(height-3),
		table.WithStyles(ts),
	)

	tb.Focus()

	return Model{
		city:   city,
		banner: banner,
		table:  tb,
		help:   help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(cmds.GetForecast(m.city))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.ForecastMsg:
		columns := utils.ForecastToColumns(msg.ForecastResponse)
		rows := utils.ForecastToRows(msg.ForecastResponse)

		m.table.SetColumns(columns)
		m.table.SetRows(rows)

		return m, nil

	case tea.KeyMsg:
		if key.Matches(msg, keymap.Quit) {
			return m, tea.Quit
		}

		if key.Matches(msg, keymap.Search) {
			utils.ClearDefaultCity()
			return m, cmds.GoToCityQuery()
		}
	}

	return m, nil
}

func (m Model) View() string {
	out := ""
	out += m.banner
	out += m.table.View()
	out += "\n"
	out += m.help.View(keymap)
	return out
}
