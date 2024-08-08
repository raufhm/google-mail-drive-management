package tuiRepo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Define the filter struct
type filter struct {
	name        string
	placeholder string
	prefix      string
}

func (f filter) Title() string {
	return f.name
}

func (f filter) Description() string {
	return f.placeholder
}

func (f filter) FilterValue() string {
	return f.name
}

var gmailFilters = []filter{
	{"Has Attachment", "Type 'yes' for filter", "has:attachment"},
	{"Google Drive/Docs Attachment", "Type 'yes' for filter", "has:drive"},
	{"Document Attachment", "Type 'yes' for filter", "has:document"},
	{"Filename", "Enter filename or file type", "filename:"},
	{"Category", "Enter category", "category:"},
	{"Size", "Enter size in bytes", "size:"},
	{"Newer than", "Enter duration (e.g., 2d, 3m, 1y)", "newer_than:"},
	{"In folder", "Enter folder name", "in:"},
}

type model struct {
	currentFilter int
	inputs        []textinput.Model
	finished      bool
}

func initialGmailModel() model {
	inputs := make([]textinput.Model, len(gmailFilters))
	for i := range inputs {
		ti := textinput.New()
		ti.Placeholder = gmailFilters[i].placeholder
		ti.CharLimit = 256
		ti.Width = 30
		if i == 0 {
			ti.Focus()
		}
		inputs[i] = ti
	}

	return model{
		currentFilter: 0,
		inputs:        inputs,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.currentFilter = len(gmailFilters) - 1
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

func (m *model) advance() {
	m.inputs[m.currentFilter].Blur()
	m.currentFilter++
	if m.currentFilter >= len(gmailFilters) {
		m.finished = true
	} else {
		m.inputs[m.currentFilter].Focus()
	}
}

func (m model) View() string {
	s := "Provide input for the filter and press Enter to move to the next filter.\nPress Ctrl+S to save and quit.\nPress Ctrl+N to skip a filter.\n\n"

	if m.finished {
		s += "Filters completed. Constructing query...\n"
		return s
	}

	s += fmt.Sprintf("Filter %d/%d: %s\n", m.currentFilter+1, len(gmailFilters), gmailFilters[m.currentFilter].name)
	s += m.inputs[m.currentFilter].View() + "\n"

	return s
}

func GetGmailQueryInput() (string, error) {
	p := tea.NewProgram(initialGmailModel())
	m, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("error running program: %w", err)
	}

	// Ensure model is of type model
	selectedModel, ok := m.(model)
	if !ok {
		return "", fmt.Errorf("type assertion failed")
	}

	var queryParts []string
	for i, input := range selectedModel.inputs {
		value := input.Value()
		if value != "" {
			if i < len(gmailFilters) {
				if gmailFilters[i].prefix == "\"" {
					queryParts = append(queryParts, fmt.Sprintf("\"%s\"", value))
				} else if value == "yes" {
					queryParts = append(queryParts, gmailFilters[i].prefix)
				} else {
					queryParts = append(queryParts, gmailFilters[i].prefix+value)
				}
			}
		}
	}

	// If no filters have been applied, use the default query
	if len(queryParts) == 0 {
		return "in:inbox category:primary", nil
	}

	return strings.Join(queryParts, " "), nil
}
