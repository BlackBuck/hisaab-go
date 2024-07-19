package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//constants

type state int

const (
	createState state = iota
	displayState	
)


// style
var (
	centerStyle = lipgloss.NewStyle().Align(lipgloss.Center).Height(8).Width(20)
)



type Component struct {
	focused state
	form    tea.Model
	list    tea.Model
}

func (c Component) Init() tea.Cmd {
	return nil
}

func (c Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.form.Update(msg)
		c.list.Update(msg)		
	case tea.KeyMsg:
		switch msg.String() {
		case "n":
			c.Update(tea.WindowSizeMsg{})
			c.focused = createState
		case "ctrl+c", "q":
			return c, tea.Quit
		case "e":
			
		}
	case Beneficiary, Expense:
		c.Update(tea.WindowSizeMsg{})
		c.list.Update(msg)
		c.focused = displayState
	}

	if c.focused == createState {
		c.form, cmd = c.form.Update(msg)
	} else {
		c.list, cmd = c.list.Update(msg)
	}
	return c, cmd
}

func (c Component) View() string {
	if c.focused == createState {
		return boxStyle.Render(centerStyle.Render(c.form.View()))
	} else {
		return boxStyle.Render(centerStyle.Render(c.list.View()))
	}
}

func NewBeneficiariesComponent(width, height int) Component {
	beneficiariesComponent := Component{}
	beneficiariesComponent.form = NewBeneficiaryForm(width, height)
	beneficiariesComponent.list = InitBeneficiaries([]Beneficiary{
		{name: "Lala", kamai: -1},
		{name: "Utkarsh", kamai: -1},
	}, width, height)

	beneficiariesComponent.focused = displayState
	return beneficiariesComponent
}

func NewExpensesComponent(width, height int) Component {
	expensesComponent := Component{}
	expensesComponent.form = NewExpensesForm(width, height)
	expensesComponent.list = InitLatestExpenses([]Expense{
		{beneficiary: "Juice", expenseType: "kharcha", amount: 50},
		{beneficiary: "Lala", expenseType: "kharcha", amount: 20},
	}, width, height)
	expensesComponent.focused = displayState
	return expensesComponent
}