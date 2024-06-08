package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/danielronalds/tgzi/tgzilib"
	"github.com/danielronalds/tgzi/tgzitui"
)

func main() {
	files, err := tgzilib.GetFiles(".", false)

	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(tgzitui.NewTuiModel(files))
	model, err := p.Run()

	if err != nil {
		panic(err)
	}

	tuiModel := model.(tgzitui.TuiModel)

	for _, file := range tuiModel.SelectedFiles {
		fmt.Println(file)
	}
}
