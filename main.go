package main

import (
	"github.com/danielronalds/tgzi/tgzilib"
	"github.com/danielronalds/tgzi/tgzitui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	files, err := tgzilib.GetFiles(".", false)

	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(tgzitui.NewTuiModel(files))
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
