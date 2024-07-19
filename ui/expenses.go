package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Expense struct {
	beneficiary string
	amount      int
	expenseType string
}

func (e Expense) FilterValue() string {
	return fmt.Sprintf("%v %v %v", e.beneficiary, e.expenseType, e.amount)
}

func (e Expense) Title() string {
	return e.beneficiary
}

func (e Expense) Description() string {
	return fmt.Sprintf("%v %v", e.expenseType, e.amount)
}

func NewExpense(beneficiary, expenseType string, amount int) Expense {
	return Expense{beneficiary: beneficiary, expenseType: expenseType, amount: amount}
}

type LatestExpenses struct {
	list list.Model
}

func (e LatestExpenses) Init() tea.Cmd {
	return nil
}

func (e LatestExpenses) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		e.list.Update(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return e, tea.Quit
		case "v":
			return e, e.GetSelectedItemInfo
		}
	case Expense:
		e.list.InsertItem(len(e.list.Items()), msg)
	}

	e.list, cmd = e.list.Update(msg)
	return e, cmd
}

func (e LatestExpenses) View() string {
	return e.list.View()
}

func InitLatestExpenses(expenses []Expense, width, height int) LatestExpenses {
	d := list.NewDefaultDelegate()
	c := lipgloss.Color("69")
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(c)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy() // copy the same to description

	res := LatestExpenses{list: list.New([]list.Item{}, d, width, height)}
	res.list.Title = "Expenses"
	res.list.SetFilteringEnabled(false) // may change later
	res.list.SetShowHelp(false)
	res.list.SetShowStatusBar(false)
	for _, beneficiary := range expenses {
		res.list.InsertItem(0, beneficiary)
	}

	return res
}

func (e Expense) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v", e.beneficiary, e.expenseType, e.amount);
}

func (m *LatestExpenses) GetSelectedItemInfo() tea.Msg {
	return m.list.SelectedItem().(Expense).GetInfo()
}