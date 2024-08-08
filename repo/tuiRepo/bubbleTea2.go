package tuiRepo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Define Google Drive specific filters
var driveFilters = []filter{
	{"Filename", "Enter filename", "name contains "},
	{"MIME Type", "Enter MIME type", "mimeType contains "},
	{"Modified Time", "Enter modified time (e.g., '>2021-01-01')", "modifiedTime > "},
	{"Size", "Enter size in bytes (e.g., '>1000')", "size > "},
	{"Full Text", "Enter text to search within files", "fullText contains "},
	//{"In Trash", "Type 'true' or 'false'", "trashed = "},
}

type driveModel struct {
	currentFilter int
	inputs        []textinput.Model
	finished      bool
}

func initialDriveModel() driveModel {
	inputs := make([]textinput.Model, len(driveFilters))
	for i := range inputs {
		ti := textinput.New()
		ti.Placeholder = driveFilters[i].placeholder
		ti.CharLimit = 256
		ti.Width = 30
		if i == 0 {
			ti.Focus()
		}
		inputs[i] = ti
	}

	return driveModel{
		currentFilter: 0,
		inputs:        inputs,
	}
}

func (m driveModel) Init() tea.Cmd {
	return nil
}

func (m driveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.finished {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			m.advance()
		case "tab":
			m.advance()
		case "shift+tab":
			m.inputs[m.currentFilter].Blur()
			m.currentFilter--
			if m.currentFilter < 0 {
				m.currentFilter = len(driveFilters) - 1
			}
			m.inputs[m.currentFilter].Focus()
		case "ctrl+s":
			m.finished = true
		case "ctrl+n":
			m.advance()
		}
	}

	var cmd tea.Cmd
	if m.currentFilter < len(m.inputs) {
		m.inputs[m.currentFilter], cmd = m.inputs[m.currentFilter].Update(msg)
	}
	return m, cmd
}

func (m *driveModel) advance() {
	m.inputs[m.currentFilter].Blur()
	m.currentFilter++
	if m.currentFilter >= len(driveFilters) {
		m.finished = true
	} else {
		m.inputs[m.currentFilter].Focus()
	}
}

func (m driveModel) View() string {
	s := "Provide input for the filter and press Enter to move to the next filter.\nPress Ctrl+S to save and quit.\nPress Ctrl+N to skip a filter.\n\n"

	if m.finished {
		s += "Filters completed. Constructing query...\n"
		return s
	}

	s += fmt.Sprintf("Filter %d/%d: %s\n", m.currentFilter+1, len(driveFilters), driveFilters[m.currentFilter].name)
	s += m.inputs[m.currentFilter].View() + "\n"

	return s
}

func GetDriveQueryInput() (string, error) {
	p := tea.NewProgram(initialDriveModel())
	m, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("error running program: %w", err)
	}

	// Ensure model is of type driveModel
	selectedModel, ok := m.(driveModel)
	if !ok {
		return "", fmt.Errorf("type assertion failed")
	}

	var queryParts []string
	for i, input := range selectedModel.inputs {
		value := input.Value()
		if value != "" {
			if i < len(driveFilters) {
				queryParts = append(queryParts, driveFilters[i].prefix+"\""+value+"\"")
			}
		}
	}

	// Join all parts of the query with ' and ' for Google Drive API query
	return strings.Join(queryParts, " and "), nil
}
