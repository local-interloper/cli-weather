package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/models/root"
)

func main() {
	app := tea.NewProgram(root.New())

	app.Run()
}
