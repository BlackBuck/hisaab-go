package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Link struct {
	title string
}

func (l Link) FilterValue() string {
	return l.title
}

func (l Link) Title() string {
	return l.title
}

func (l Link) Description() string {
	return ""	
}

type SideBar struct {
	links list.Model
	choice string
}

func (s SideBar) Init()	tea.Cmd {
	return nil
}

func (s SideBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.links.SetWidth(msg.Width/4)
		s.links.SetHeight(msg.Height-10)
		return s, nil	
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		case "tab":
			s.choice = s.links.SelectedItem().FilterValue()
			return s, nil
		}
	}

	var cmd tea.Cmd
	s.links, cmd = s.links.Update(msg)
	return s, cmd
}

func (s SideBar) View() string {
	return boxStyle.Render(s.links.View())
}

func NewSideBar(links []string) SideBar{
	d := list.NewDefaultDelegate()
	c := lipgloss.Color("69")
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(c)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy() // copy the same to description

	res := SideBar{links: list.New([]list.Item{}, d, 100, 100)}
	res.links.Title = "HisaabKitaab"

	//don't filter values, don't show help
	res.links.SetFilteringEnabled(false)
	res.links.SetShowHelp(false)
	res.links.SetShowStatusBar(false)
	for _, link := range links {
		res.links.InsertItem(len(res.links.Items()), Link{title: link})
	}

	return res
}