package utils

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/local-interloper/cli-weather/app/types"
)

func CitiesToItems(cities []types.City) []list.Item {
	items := []list.Item{}

	for _, city := range cities {
		items = append(items, city)
	}

	return items
}
