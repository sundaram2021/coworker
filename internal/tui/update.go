package tui

import (
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sundaram2021/coworker/internal/config"
	"github.com/sundaram2021/coworker/internal/utils"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.quit = true
			return m, tea.Quit

		case tea.KeyEnter:
			switch m.state {
			case viewInput:
				if !m.apiKeyExists {
					m.state = viewAPIModal
					m.apiKeyInput.Focus()
					m.textInput.Blur()
					return m, nil
				}
				if m.textInput.Value() != "" || len(m.attachedFiles) > 0 {
					m.textInput.SetValue("")
					m.attachedFiles = []string{}
				}

			case viewFilePicker:
				if m.cursor < len(m.files) {
					m.attachedFiles = append(m.attachedFiles, m.files[m.cursor].Name())
					m.state = viewInput
					m.textInput.Focus()
				}

			case viewAPIModal:
				if m.apiKeyInput.Value() != "" {
					m.state = viewSaving
					m.apiKeyInput.Blur()
					return m, tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
						err := os.MkdirAll(filepath.Dir(config.ApiKeyPath), 0755)
						if err != nil {
							return SaveKeyErrMsg{Err: err}
						}
						err = os.WriteFile(config.ApiKeyPath, []byte(m.apiKeyInput.Value()), 0600)
						if err != nil {
							return SaveKeyErrMsg{Err: err}
						}
						return SaveKeySuccessMsg{}
					})
				}
			}

		case tea.KeyEsc:
			if m.state == viewFilePicker {
				m.state = viewInput
				m.textInput.Focus()
			} else if m.state == viewAPIModal {
				m.state = viewInput
				m.textInput.Focus()
			}

		case tea.KeyCtrlF:
			if m.state == viewInput {
				m.state = viewFilePicker
				m.textInput.Blur()
				files, err := utils.LoadFiles(".")
				if err == nil {
					m.files = files
					m.cursor = 0
				} else {
					m.files = []os.DirEntry{}
				}
			}

		case tea.KeyCtrlG:
			if m.state == viewInput {
				m.state = viewAPIModal
				m.apiKeyInput.Focus()
				m.textInput.Blur()
			}

		case tea.KeyUp:
			if m.state == viewFilePicker && m.cursor > 0 {
				m.cursor--
			}

		case tea.KeyDown:
			if m.state == viewFilePicker && m.cursor < len(m.files)-1 {
				m.cursor++
			}

		case tea.KeyRunes:
			for _, r := range msg.Runes {
				if m.state == viewFilePicker {
					switch r {
					case 'k':
						if m.cursor > 0 {
							m.cursor--
						}
					case 'j':
						if m.cursor < len(m.files)-1 {
							m.cursor++
						}
					}
				}
			}
		}

		if m.state == viewInput {
			m.textInput, cmd = m.textInput.Update(msg)
		} else if m.state == viewAPIModal {
			m.apiKeyInput, cmd = m.apiKeyInput.Update(msg)
		}

	case SaveKeySuccessMsg:
		m.apiKeyExists = true
		m.state = viewInput
		m.textInput.Focus()
		m.apiKeyInput.SetValue("")
		return m, nil

	case SaveKeyErrMsg:
		m.err = msg.Err
		m.state = viewInput
		m.textInput.Focus()
		return m, nil

	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft && m.state == viewInput {
			if msg.X < 4 && msg.Y > 6 && msg.Y < 10 {
				m.state = viewFilePicker
				m.textInput.Blur()
				files, err := utils.LoadFiles(".")
				if err == nil {
					m.files = files
					m.cursor = 0
				} else {
					m.files = []os.DirEntry{}
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case spinner.TickMsg:
		if m.state == viewSaving {
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}

	return m, cmd
}
