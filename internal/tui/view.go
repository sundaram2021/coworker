package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sundaram2021/coworker/internal/config"
	"github.com/sundaram2021/coworker/internal/styles"
)

func (m model) View() string {

	// Title Art
	title := styles.BigTitleStyle.Render(styles.TitleArt)

	// Subtitle
	subtitle := styles.SubtitleStyle.Render("Your Coworker for your tasks")

	centeredTitle := lipgloss.Place(config.ContentWidth, lipgloss.Height(title), lipgloss.Center, lipgloss.Center, title)
	centeredSubtitle := lipgloss.Place(config.ContentWidth, lipgloss.Height(subtitle), lipgloss.Center, lipgloss.Center, subtitle)

	// Input Area
	// TODO: in icons we need better implementation , don't like the current one
	icon := "📎"   
	if len(m.attachedFiles) > 0 {
		icon = "📄"
	}

	inputRow := strings.Builder{}
	inputRow.WriteString(styles.FileAttachmentStyle.Render(icon + " "))
	inputRow.WriteString(m.textInput.View())

	inputContainer := styles.InputContainerStyle.Render(inputRow.String())

	// Attached Files List
	var filesDisplay string
	if len(m.attachedFiles) > 0 {
		var fileList []string
		for _, f := range m.attachedFiles {
			fileList = append(fileList, styles.SelectedFileStyle.Render(f))
		}
		filesDisplay = strings.Join(fileList, " ")
	} else {
		filesDisplay = ""
	}

	centeredFiles := lipgloss.Place(config.ContentWidth, lipgloss.Height(filesDisplay), lipgloss.Center, lipgloss.Center, filesDisplay)

	// Footer / Help
	footerText := styles.HelpStyle.Render("Ctrl+F: Attach File  |  Ctrl+G: Update API Key  |  Enter: Send")
	centeredFooter := lipgloss.Place(config.ContentWidth, lipgloss.Height(footerText), lipgloss.Center, lipgloss.Center, footerText)

	tipText := styles.TipStyle.Render("Tip: Use coworker run -f file.ts to attach files via CLI")
	centeredTip := lipgloss.Place(config.ContentWidth, lipgloss.Height(tipText), lipgloss.Center, lipgloss.Center, tipText)

	mainLayout := lipgloss.JoinVertical(
		lipgloss.Left,
		centeredTitle,
		centeredSubtitle,
		"\n", // Spacer
		inputContainer,
		"\n",
		centeredFiles,
		"\n",
		centeredFooter,
		centeredTip,
	)

	// Handle Modals (Overlays)

	if m.state == viewAPIModal {
		// Render Modal Content
		modalHeader := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#88C0D0")).
			Bold(true).
			MarginBottom(1).
			Render("Gemini API Key Required")

		modalInput := m.apiKeyInput.View()

		modalFooter := styles.HelpStyle.Render("Enter: Save  |  Esc: Cancel")

		modalBox := styles.ModalContentStyle.Render(
			fmt.Sprintf("%s\n\n%s\n\n%s", modalHeader, modalInput, modalFooter),
		)

		// Center the modal box in the middle of the screen
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			modalBox,
		)
	} else if m.state == viewSaving {
		loadingText := fmt.Sprintf("Saving API Key... %s", m.spinner.View())
		modalBox := styles.ModalContentStyle.Align(lipgloss.Center, lipgloss.Center).Render(loadingText)

		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			modalBox,
		)
	} else if m.state == viewFilePicker {
		// File Picker Popup
		popupStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#EBCB8B")).
			Background(lipgloss.Color("#2E3440")).
			Padding(1).
			Width(60).
			Height(15)

		header := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EBCB8B")).
			Bold(true).
			Background(lipgloss.Color("#2E3440")).
			Width(56).
			Render("Select a File to Attach:")

		var filesContent string
		if len(m.files) == 0 {
			filesContent = lipgloss.NewStyle().Faint(true).Render("No files found in current directory.")
		} else {
			var lines []string
			for i, file := range m.files {
				cursor := " "
				if i == m.cursor {
					cursor = ">"
				}
				line := fmt.Sprintf("%s %s", cursor, file.Name())
				if i == m.cursor {
					line = lipgloss.NewStyle().
						Foreground(lipgloss.Color("#88C0D0")).
						Bold(true).
						Background(lipgloss.Color("#2E3440")).
						Width(56).
						Render(line)
				}
				lines = append(lines, line)
			}
			filesContent = strings.Join(lines, "\n")
		}

		helpText := "\n" + styles.HelpStyle.
			Background(lipgloss.Color("#2E3440")).
			Width(56).
			Align(lipgloss.Center).
			Render("Enter: Select  |  Esc: Cancel  |  ↑/↓: Navigate")

		popupContent := lipgloss.JoinVertical(lipgloss.Left, header, filesContent, helpText)
		popupBox := popupStyle.Render(popupContent)

		// Center the file picker in the middle of the screen
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			popupBox,
		)
	}

	// Place the main layout in the top-center of the screen
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		mainLayout,
	)
}
