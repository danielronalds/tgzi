package tgzitui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const filesPerPage int = 10

// The data model for the TUI containing the state of the interface
type TuiModel struct {
	files    []string
	cursor   int
	selected map[int]struct{}
	help     bool
	min      int
	max      int
	// The exported list of files that were selected. Populated on exit
	SelectedFiles []string
}

// Creates a new tui for selecting from the given files
func NewTuiModel(files []string) TuiModel {
	return TuiModel{
		files:         files,
		selected:      make(map[int]struct{}),
		SelectedFiles: make([]string, 0),
		min:           0,
		max:           10,
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
				if m.cursor < m.min {
					m.min -= filesPerPage;
					m.max -= filesPerPage;
				}
			}

		case "down", "j":
			if m.cursor < len(m.files)-1 {
				m.cursor++
				if m.cursor > m.max {
					m.min += filesPerPage;
					m.max += filesPerPage;
				}
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

		case "?":
			m.help = !m.help
		}

	}

	return m, nil
}

func (m TuiModel) View() string {
	if m.help {
		s := "Key        Action"
		s += "\nup/k       navigate up the list"
		s += "\ndown/j     navigate down the list"
		s += "\nspace      Select a file"
		s += "\nenter      Compress files"
		s += "\n?          toggle help menu"
		return s
	}

	s := "Select files to compress"

	currPage := (m.cursor / filesPerPage) + 1
	maxPages := (len(m.files) / filesPerPage) + 1
	if maxPages > 1 {
		s += fmt.Sprintf(" (%v/%v)", currPage,  maxPages)
	}

	for i, file := range m.files {
		if i < m.min || i > m.max {
			continue;
		}

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("\n%s [%s] %s", cursor, checked, file)
	}

	s += "\nPress ? for help"
	return s
}

// Runs a one line prompt for the user to type the filename of the archive into
func GetArchiveName() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Archive name (Leave blank for default)\n> ")

	archive, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	archive = strings.Replace(archive, "\n", "", -1)

	return archive, nil
}
