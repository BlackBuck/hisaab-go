package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ExpensesInfo struct {
	boxes []BoxModel
}

func (e ExpensesInfo) Init() tea.Cmd {
	return nil
}

func (e ExpensesInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for _, box := range e.boxes {
			box.Update(msg)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return e, tea.Quit
		}
	}

	return e, nil	
}

func (e ExpensesInfo) View() string {
	s := []string{}
	for _, box := range e.boxes {
		s = append(s, box.View())
	}
	return lipgloss.NewStyle().MarginTop(2).Render(lipgloss.JoinHorizontal(lipgloss.Center, s...))
}

func InitGrid(titles, descriptions []string) ExpensesInfo {
	expenses := ExpensesInfo{boxes: []BoxModel{}}
	for i := 0;i < len(titles);i++ {
		expenses.boxes = append(expenses.boxes, BoxModel{title: titles[i], description: descriptions[i]})
	}
	return expenses
}