package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	infoStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("69")).MaxHeight(5)
)

type InfoModel struct {
	info string
}

func (i InfoModel) Init() tea.Cmd {
	return nil	
}

func (i InfoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		infoStyle.Width(msg.Width)
	}
	return i, cmd
}

func (i InfoModel) View() string {
	return infoStyle.Render(i.info)
}

func DefaultInfoModel() InfoModel {
	return InfoModel{info: "Press v on a beneficiary or expense to get details."}
}

func (i *InfoModel) SetInfo(info string) {
	i.info = info
}

func InfoModelWithData(info string) InfoModel {
	return InfoModel{info: info}
}