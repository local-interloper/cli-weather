package main

import (
	"flag"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		println("Usage: weather cityName")
		return
	}

	app := tea.NewProgram(app.Make(args[0]))
	app.Run()
}
