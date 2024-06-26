package main

import (
	"fmt"
	"os"

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

	if len(tuiModel.SelectedFiles) == 0 {
		return;
	}

	archive, _ := tgzitui.GetArchiveName()
	archive = tgzilib.NormaliseArchiveName(archive)

	err = tgzilib.CompressFiles(archive, tuiModel.SelectedFiles)
	if err != nil {
		fmt.Printf("Recieved an error attempting to compress files: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Created %s\n", archive)
}
