package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	boxStyle = lipgloss.NewStyle().Padding(1, 2).Margin(1, 1).Border(lipgloss.RoundedBorder()).AlignVertical(lipgloss.Center)
	titleStyle = lipgloss.NewStyle().Margin(1, 1).Padding(0, 1).Background(lipgloss.Color("69"))
	descriptionStyle = lipgloss.NewStyle().Margin(1, 1).Padding(0, 1).Background(lipgloss.Color("240"))
)

type BoxModel struct {
	title, description string
}

func (m BoxModel) Init() tea.Cmd {
	return nil	
}

func (m BoxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		boxStyle.Width(msg.Width).Height(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		
		}
	}
	return m, nil
}

func (m BoxModel) View() string {
	s := titleStyle.Render(m.title)
	s = lipgloss.JoinVertical(lipgloss.Top, s, descriptionStyle.Render(m.description))
	return boxStyle.Render(s)
}

func InitModel(title, description string) BoxModel {
	return BoxModel{title: title, description: description}
}