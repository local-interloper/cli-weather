package models

import (
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

type forecastModelKeymapType struct {
	Search key.Binding
	Quit   key.Binding
}

func (k forecastModelKeymapType) ShortHelp() []key.Binding {
	return []key.Binding{k.Search, k.Quit}
}

func (k forecastModelKeymapType) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Search, k.Quit}}
}

var forecastModelKeymap forecastModelKeymapType = forecastModelKeymapType{
	Search: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("S", "Search"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("Ctrl+C", "Quit"),
	),
}

type ForecastModel struct {
	city  types.City
	table table.Model
	help  help.Model
}

func MakeForecastModel(city types.City) ForecastModel {
	width, height := utils.TermSize()

	ts := table.DefaultStyles()
	ts.Header = ts.Header.Background(lipgloss.Color("#8888ff"))
	ts.Selected = ts.Selected.Foreground(lipgloss.NoColor{}).Bold(false)

	tb := table.New(
		table.WithWidth(width),
		table.WithHeight(height-1),
		table.WithStyles(ts),
	)

	tb.Focus()

	return ForecastModel{
		city:  city,
		table: tb,
		help:  help.New(),
	}
}

func (m ForecastModel) Init() tea.Cmd {
	return tea.Batch(cmds.GetForecast(m.city))
}

func (m ForecastModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.ForecastMsg:
		columns := utils.ForecastToColumns(msg.ForecastResponse)
		rows := utils.ForecastToRows(msg.ForecastResponse)

		m.table.SetColumns(columns)
		m.table.SetRows(rows)
		return m, nil

	case tea.KeyMsg:
		if key.Matches(msg, forecastModelKeymap.Quit) {
			return m, tea.Quit
		}

		if key.Matches(msg, forecastModelKeymap.Search) {
			utils.ClearDefaultCity()
			return m, cmds.SwitchModel(MakeSetupModel())
		}
	}

	return m, nil
}

func (m ForecastModel) View() string {
	out := ""
	out += m.table.View()
	out += m.help.View(forecastModelKeymap)
	return out
}
