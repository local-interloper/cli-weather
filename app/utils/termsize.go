package utils

import (
	"os"

	"github.com/charmbracelet/x/term"
)

func TermSize() (int, int) {
	width, height, err := term.GetSize(os.Stdin.Fd())

	if err != nil {
		panic("Failed to obtain temrinal size")
	}

	return width, height
}
