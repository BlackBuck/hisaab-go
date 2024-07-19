package ui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	beneficiarySelect status = iota
	amountSelect
	typeSelect
)

type ExpenseType struct {
	expenseType string
}

func (t ExpenseType) FilterValue() string {
	return t.expenseType
}

func (t ExpenseType) Title() string {
	return t.expenseType
}

func (t ExpenseType) Description() string {
	return ""
}

type ExpenseFormModel struct {
	focused     status
	beneficiary list.Model
	amount      textinput.Model
	expenseType list.Model
}

func (e ExpenseFormModel) Init() tea.Cmd {
	return nil
}

func (e ExpenseFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		e.beneficiary.Update(msg)
		e.amount.Update(msg)
		e.expenseType.Update(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if e.focused != typeSelect {
				e.focused++
			} else {
				e.Update(tea.WindowSizeMsg{})
				return e, e.NewExpense
			}
		case "backspace":
			if e.focused != beneficiarySelect {
				e.focused--
			}
		}
	}

	// e.Update(tea.WindowSizeMsg{})
	if e.focused == beneficiarySelect {
		e.beneficiary, cmd = e.beneficiary.Update(msg)
	} else if e.focused == amountSelect {
		e.amount, cmd = e.amount.Update(msg)
	} else {
		e.expenseType, cmd = e.expenseType.Update(msg)
	}

	return e, cmd
}

func (e ExpenseFormModel) View() string {
	switch e.focused {
	case beneficiarySelect:
		return lipgloss.JoinVertical(lipgloss.Top, titleStyle.Render("Select Benef"), e.beneficiary.View())
	case amountSelect:
		return lipgloss.JoinVertical(lipgloss.Top, titleStyle.Render("Enter Amount"), e.amount.View())
	case typeSelect:
		return lipgloss.JoinVertical(lipgloss.Top, titleStyle.Render("Select Type"), e.expenseType.View())
	default:
		return ""
	}
}

func (e *ExpenseFormModel) NewExpense() tea.Msg {
	amount, err := strconv.Atoi(e.amount.Value())
	if err != nil {
		e.focused = amountSelect
		return tea.KeyBackspace
	}
	return NewExpense(e.beneficiary.SelectedItem().FilterValue(), e.expenseType.SelectedItem().FilterValue(), amount)
}

func NewExpensesForm(width, height int) ExpenseFormModel {
	d := list.NewDefaultDelegate()
	c := lipgloss.Color("69")
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(c)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy() // copy the same to description
	d.Styles.NormalDesc = d.Styles.NormalDesc.Foreground(lipgloss.Color("240")).Align(lipgloss.Left)
	d.Styles.NormalTitle = d.Styles.NormalDesc.Copy()

	beneficiaries := []Beneficiary{
		{name: "Utkarsh", kharcha: -1},
		{name: "Lala", kharcha: -1},
	}

	expensesForm := ExpenseFormModel{beneficiary: list.New([]list.Item{}, d, width, height)}
	expensesForm.beneficiary.SetFilteringEnabled(false) // may change later
	expensesForm.beneficiary.SetShowHelp(false)
	expensesForm.beneficiary.SetShowStatusBar(false)
	expensesForm.beneficiary.SetShowTitle(false)
	for _, beneficiary := range beneficiaries {
		expensesForm.beneficiary.InsertItem(0, beneficiary)
	}

	//Text input for amount
	expensesForm.amount = textinput.New()
	expensesForm.amount.Placeholder = "Rupees"
	expensesForm.amount.Focus()


	//Types of expenses
	d2 := list.NewDefaultDelegate()
	d2.Styles.NormalTitle.Padding(1, 2)
	expensesForm.expenseType = list.New([]list.Item{}, d2, width, height)
	expensesForm.expenseType.InsertItem(0, ExpenseType{expenseType: "kharcha"})
	expensesForm.expenseType.InsertItem(1, ExpenseType{expenseType: "kamai"})
	expensesForm.expenseType.InsertItem(2, ExpenseType{expenseType: "udhari"})
	expensesForm.expenseType.InsertItem(3, ExpenseType{expenseType: "bakaya"})
	expensesForm.expenseType.SetShowStatusBar(false)
	expensesForm.expenseType.SetShowFilter(false)
	expensesForm.expenseType.SetShowHelp(false)
	expensesForm.expenseType.SetShowTitle(false)

	expensesForm.focused = beneficiarySelect

	return expensesForm
}
