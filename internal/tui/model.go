package tui

import (
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sundaram2021/coworker/internal/config"
	"github.com/sundaram2021/coworker/internal/styles"
)

type state int

const (
	viewInput state = iota
	viewFilePicker
	viewAPIModal
	viewSaving
)

type model struct {
	textInput     textinput.Model
	apiKeyInput   textinput.Model
	spinner       spinner.Model
	attachedFiles []string
	files         []os.DirEntry
	cursor        int
	width         int
	height        int
	state         state
	apiKeyExists  bool
	err           error
	quit          bool
}

func InitialModel() tea.Model {
	ti := textinput.New()
	ti.Placeholder = "Ask anything..."
	ti.Focus()
	ti.Width = 60 // Slightly smaller than container to fit padding
	ti.TextStyle.Foreground(lipgloss.Color("#ECEFF4"))
	ti.PlaceholderStyle.Foreground(lipgloss.Color("#4C566A"))
	ti.EchoMode = textinput.EchoNormal

	aki := textinput.New()
	aki.Placeholder = "Enter your Gemini API Key"
	aki.Width = 50
	aki.TextStyle.Foreground(lipgloss.Color("#ECEFF4"))
	aki.PlaceholderStyle.Foreground(lipgloss.Color("#4C566A"))
	aki.EchoMode = textinput.EchoPassword

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = styles.SpinnerStyle

	_, err := os.Stat(config.ApiKeyPath)
	exists := false
	if err == nil {
		exists = true
	}

	return model{
		textInput:     ti,
		apiKeyInput:   aki,
		spinner:       s,
		state:         viewInput,
		attachedFiles: []string{},
		files:         []os.DirEntry{},
		cursor:        0,
		apiKeyExists:  exists,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
