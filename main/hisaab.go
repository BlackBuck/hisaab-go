package main

import (
	"fmt"
	"os"
	"ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedBorderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	normalBorderStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	centerStyle         = lipgloss.NewStyle().AlignVertical(lipgloss.Center).Height(30).Width(30)
	mainModelStyle      = lipgloss.NewStyle().Padding(1, 1).Border(lipgloss.RoundedBorder()).MaxHeight(60)
	titleStyle			= lipgloss.NewStyle().Padding(1, 1).Background(lipgloss.Color("12")).Border(lipgloss.RoundedBorder()).Height(5).AlignVertical(lipgloss.Center)
)

type MainModel struct {
	title string
	infoBar ui.InfoModel
	expenseList tea.Model
	beneficiaryList tea.Model
	expenseGrid ui.ExpensesInfo
	chartModel tea.Model
	current int
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.beneficiaryList.Update(msg)
		m.expenseGrid.Update(msg)
		m.expenseList.Update(msg)
		m.infoBar.Update(msg)
		m.chartModel.Update(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.current = (m.current%3) + 1 //one or two
			m.infoBar = ui.DefaultInfoModel()
			return m, nil
		}
	case string:
		m.infoBar = ui.InfoModelWithData(msg)
		m.infoBar.Update(tea.WindowSizeMsg{})
	default:
		// for _, model := range m.models {
		// 	model.Update(msg)
		// }
	}

	if m.current == 1 {
		m.beneficiaryList, cmd = m.beneficiaryList.Update(msg)
	} else if m.current == 2{
		m.expenseList, cmd = m.expenseList.Update(msg)
	} else {
		m.chartModel, cmd = m.chartModel.Update(msg)
	}
	
	return m, cmd
}

func (m MainModel) View() string {
	res := []string{}
	if m.current == 1 {
		res = append(res, selectedBorderStyle.Render(centerStyle.Render(m.beneficiaryList.View())))
		res = append(res, normalBorderStyle.Render(centerStyle.Render(m.expenseList.View())))
		res = append(res, normalBorderStyle.Render(centerStyle.Render(m.chartModel.View())))
	} else if m.current == 2 {
		res = append(res, normalBorderStyle.Render(centerStyle.Render(m.beneficiaryList.View())))
		res = append(res, selectedBorderStyle.Render(centerStyle.Render(m.expenseList.View())))
		res = append(res, normalBorderStyle.Render(centerStyle.Render(m.chartModel.View())))
	} else {
		res = append(res, normalBorderStyle.Render(centerStyle.Render(m.beneficiaryList.View())))
		res = append(res, normalBorderStyle.Render(centerStyle.Render(m.expenseList.View())))
		res = append(res, selectedBorderStyle.Render(centerStyle.Render(m.chartModel.View())))
	}

	return mainModelStyle.Render(lipgloss.JoinVertical(lipgloss.Top, titleStyle.Render(m.title), m.expenseGrid.View(),lipgloss.JoinHorizontal(lipgloss.Left, res...), m.infoBar.View()))
}

func main() {
	width, height := 20, 20

	beneficiariesComponent := ui.NewBeneficiariesComponent(width, height)
	expensesComponent := ui.NewExpensesComponent(width, height)

	mainModel := MainModel{}
	mainModel.title = "HisaabKitaab"
	mainModel.beneficiaryList = beneficiariesComponent
	mainModel.expenseList = expensesComponent
	mainModel.current = 0
	mainModel.infoBar = ui.DefaultInfoModel()
	mainModel.expenseGrid = ui.InitGrid([]string{"Kharcha", "Kamai", "Udhari", "Bakaya", "Total"},
										[]string{"10", "20", "30", "40", "100"})
	mainModel.chartModel = ui.NewBarChart()

	p := tea.NewProgram(mainModel, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("An error occured in main program: ", err)
		os.Exit(1)
	}
}
