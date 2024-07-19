package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Beneficiary struct {
	name                           string
	kharcha, kamai, udhari, bakaya int
}

func (b Beneficiary) FilterValue() string {
	return b.name
}

func (b Beneficiary) Title() string {
	return b.name
}

func (b Beneficiary) Description() string {
	// here negative means that
	//the beneficiary is a part of some other list
	//other than the beneficiary list
	if b.kharcha == -1 || b.kamai == -1 || b.udhari == -1 || b.bakaya == -1 {
		return ""
	} else {
		return fmt.Sprintf("Total Hisaab--%v", b.kharcha+b.kamai+b.bakaya+b.udhari)
	}
}

type Beneficiaries struct {
	list list.Model
}

func (b Beneficiaries) Init() tea.Cmd {
	return nil
}

func (b Beneficiaries) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.Update(msg)
	case Beneficiary:
		b.list.InsertItem(len(b.list.Items()), msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "v":
			// viewing the details of the selected beneficiary
			return b, b.GetSelectedItemInfo
		}
	}

	b.list, cmd = b.list.Update(msg)
	return b, cmd
}

func (b Beneficiaries) View() string {
	return b.list.View()
}

func InitBeneficiaries(beneficiaries []Beneficiary, width, height int) Beneficiaries {
	d := list.NewDefaultDelegate()
	c := lipgloss.Color("69")
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(c)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy() // copy the same to description

	res := Beneficiaries{list: list.New([]list.Item{}, d, width, height)}
	res.list.Title = "Beneficiaries"
	res.list.SetFilteringEnabled(false) // may change later
	res.list.SetShowHelp(false)
	res.list.SetShowStatusBar(false)
	for _, beneficiary := range beneficiaries {
		res.list.InsertItem(0, beneficiary)
	}

	return res
}

func (b Beneficiary) GetInfo() string {
	if b.kharcha < 0 || b.kamai < 0 || b.udhari < 0 || b.bakaya < 0 {
		// no hisaab so far
		return fmt.Sprintf("%v NO HISAAB", b.name)
	} else {
		return fmt.Sprintf("%v Kharcha - %v Kamai - %v Udhari - %v Bakaya - %v", b.name, b.kharcha, b.kamai, b.udhari, b.bakaya)
	}
}

func (m *Beneficiaries) GetSelectedItemInfo() tea.Msg {
	return m.list.SelectedItem().(Beneficiary).GetInfo()
}