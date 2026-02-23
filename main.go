package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/models"
)

func main() {
	app := tea.NewProgram(models.MakeRootModel())

	app.Run()
}
