package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sundaram2021/coworker/internal/config"
)

// --- ASCII Art for "COWORKER" --- , may be we change this static implementation to something more robust
const TitleArt = `
  _____    ____  __          __   ____   _____   _  __  ______  _____  
 / ____|  / __ \ \ \        / /  / __ \ |  __ \ | |/ / |  ____||  __ \ 
| |      | |  | | \ \  /\  / /  | |  | || |__) || ' /  | |__   | |__) |
| |      | |  | |  \ \/  \/ /   | |  | ||  _  / |  <   |  __|  |  _  / 
| |____  | |__| |   \  /\  /    | |__| || | \ \ | . \  | |____ | | \ \ 
 \_____|  \____/     \/  \/      \____/ |_|  \_\|_|\_\ |______||_|  \_\
`

// --- Styles ---
var (
	BigTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8FBCBB")).
			Bold(true)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#88C0D0")).
			Faint(true)

	InputContainerStyle = lipgloss.NewStyle().
				Width(config.ContentWidth).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#4C566A")).
				Padding(1).
				Align(lipgloss.Center)

	FileAttachmentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#EBCB8B")).
				PaddingRight(1)

	SelectedFileStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#81A1C1")).
		Padding(0, 1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#616E88")).
			MarginTop(1).
			Align(lipgloss.Center)

	TipStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4C566A")).
			MarginTop(2).
			Align(lipgloss.Center)

	ModalContentStyle = lipgloss.NewStyle().
				Width(60).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#88C0D0")).
				Background(lipgloss.Color("#2E3440")).
				Padding(2).
				Align(lipgloss.Center)

	SpinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8FBCBB"))

	ErrorMsgStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#BF616A")).
			MarginTop(1)
)
