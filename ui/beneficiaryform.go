package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)



func NewBeneficiary(name string) Beneficiary {
	beneficiary := Beneficiary{name: name, kharcha: -1}
	return beneficiary
}

// FORM MODEL
type BeneficiaryFormModel struct {
	name textinput.Model	
}

func NewForm() *BeneficiaryFormModel {
	form := BeneficiaryFormModel{name: textinput.New()}
	form.name.Focus()
	return &form
}

func NewFormWithData(name string) *BeneficiaryFormModel {
	form := BeneficiaryFormModel{name: textinput.New()}
	form.name.SetValue(name)
	form.name.Focus()
	return &form
}

func (m BeneficiaryFormModel) Init() tea.Cmd {
	return nil	
}

func (m BeneficiaryFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.name.Update(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.Update(tea.KeyTab)
			return m, m.NewBeneficiary
		}
	}
	
	m.name, cmd = m.name.Update(msg)
	return m, cmd 
}

func (m BeneficiaryFormModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Top, titleStyle.Render("New Beneficiary"), m.name.View())
}

func (m *BeneficiaryFormModel) NewBeneficiary() tea.Msg {
	beneficiary := NewBeneficiary(m.name.Value())
	return beneficiary
}

func NewBeneficiaryForm(width, height int) BeneficiaryFormModel {
	beneficiaryForm := BeneficiaryFormModel{}
	beneficiaryForm.name = textinput.New()
	beneficiaryForm.name.Focus()

	return beneficiaryForm
}