package tgzitui

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sort.Strings(files)

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
			if !m.help && m.cursor > 0 {
				m.cursor--
				if m.cursor < m.min {
					m.min -= filesPerPage
					m.max -= filesPerPage
				}
			}

		case "down", "j":
			if !m.help && m.cursor < len(m.files)-1 {
				m.cursor++
				if m.cursor > m.max {
					m.min += filesPerPage
					m.max += filesPerPage
				}
			}

		case " ":
			if !m.help {
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}

		case "A":
			if !m.help {
				for i := range m.files {
					m.selected[i] = struct{}{}
				}
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
		s := `┌────────┬─────────────────────────┐
│ Key    │ Action                  │
├────────┼─────────────────────────┤
│ up/k   │  Navigate up the list   │
│ down/j │  Navigate down the list │
│ space  │  Select a file          │
│ A      │  Select all files       │
│ enter  │  Compress files         │
│ ?      │  Toggle help menu       │
└────────┴─────────────────────────┘

tgzi v1.0.0
`
		return s
	}

	s := "Select files to compress"

	currPage := (m.cursor / filesPerPage) + 1
	maxPages := (len(m.files) / filesPerPage) + 1
	if maxPages > 1 {
		s += fmt.Sprintf(" (%v/%v)", currPage, maxPages)
	}

	filesPrinted := 0

	for i, file := range m.files {
		if i < m.min || i > m.max {
			continue
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

		filesPrinted += 1
	}

	for len(m.files) > filesPerPage && ((filesPerPage+1)-filesPrinted) > 0 {
		s += "\n"
		filesPrinted += 1
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
