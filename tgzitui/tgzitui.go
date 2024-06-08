package tgzitui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type TuiModel struct {
	files    []string
	cursor   int
	selected map[int]struct{}
	// The exported list of files that were selected. Populated on exit
	SelectedFiles []string
}

func NewTuiModel(files []string) TuiModel {
	return TuiModel{
		files:         files,
		selected:      make(map[int]struct{}),
		SelectedFiles: make([]string, 0),
	}
}

func (m TuiModel) Init() tea.Cmd {
	return nil
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.files)-1 {
				m.cursor++
			}

		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		case "enter":
			for i, file := range m.files {
				if _, ok := m.selected[i]; ok {
					m.SelectedFiles = append(m.SelectedFiles, file)
				}
			}

			return m, tea.Quit
		}

	}

	return m, nil
}

func (m TuiModel) View() string {
	s := "Select files to compress\n"

	for i, file := range m.files {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, file)
	}

	s += "\nPress Enter to compress"

	return s
}
