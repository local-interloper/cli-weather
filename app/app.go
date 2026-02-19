package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/local-interloper/cli-weather/app/cmds"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/utils"
)

type AppModel struct {
	spin  spinner.Model
	table table.Model
	ready bool
	city  string
}

func Make(city string) AppModel {
	spin := spinner.New()
	spin.Spinner = spinner.Line

	ts := table.DefaultStyles()
	ts.Header = ts.Header.Background(lipgloss.Color("#8888ff"))
	ts.Selected = ts.Selected.Foreground(lipgloss.NoColor{}).Bold(false)

	return AppModel{
		spin: spin,
		city: city,
		table: table.New(
			table.WithHeight(10),
			table.WithStyles(ts),
		),
		ready: false,
	}
}

func (a AppModel) Init() tea.Cmd {
	return tea.Batch(a.spin.Tick, cmds.GetForecast(a.city))
}

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.ForecastMsg:
		a.table.SetColumns(utils.ColumnsFromForecastMsg(msg))
		a.table.SetRows(utils.RowsFromForecastMsg(msg))
		a.ready = true
		return a, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}

	case spinner.TickMsg:
		if a.ready {
			return a, nil
		}

		var cmd tea.Cmd
		a.spin, cmd = a.spin.Update(msg)
		return a, cmd
	}

	return a, nil
}

func (a AppModel) View() string {
	if !a.ready {
		return fmt.Sprintf("Fetching weather %s\n", a.spin.View())
	}

	return a.table.View()
}
