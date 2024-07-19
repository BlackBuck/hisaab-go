package ui

import (
	"fmt"

	"github.com/NimbleMarkets/ntcharts/barchart"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type ChartModel struct {
	chart 		barchart.Model
	max, mean   float64
}

func (c ChartModel) Init() tea.Cmd {
	return nil	
}

func (c ChartModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	c.chart.Draw()
	return c, nil	
}

func (c ChartModel) View() string {
	stats := lipgloss.JoinHorizontal(lipgloss.Left, 
				boxStyle.Height(4).Width(10).UnsetMarginTop().Render(fmt.Sprintf("Max - %v", c.max)),
				boxStyle.Height(4).Width(10).UnsetMarginTop().Render(fmt.Sprintf("Mean - %v", c.mean)))
	chart := boxStyle.Render(centerStyle.Render(lipgloss.JoinVertical(lipgloss.Top, titleStyle.Render("Monthly"), c.chart.View())))

	return lipgloss.JoinVertical(lipgloss.Top, chart, stats)
}


func NewBarChart() ChartModel{
	v1 := barchart.BarData{
		Label: "Jun",
		Values: []barchart.BarValue{
			{Name: "June", Value: 21.2},
		},
	}
    v2 := barchart.BarData{
		Label: "Jul",
		Values: []barchart.BarValue{
			{Name: "July", Value: 19.2},
		},
	}
    v3 := barchart.BarData{
		Label: "Aug",
		Values: []barchart.BarValue{
			{Name: "Aug", Value: 11.2},
		},
	}
    
    bc := barchart.New(14, 10, 
                    barchart.WithDataSet([]barchart.BarData{v1, v2, v3}),
                    barchart.WithBarGap(1),
                    barchart.WithMaxValue(22),
                    barchart.WithBarWidth(1))
    
    
	return ChartModel{chart: bc, max: 22.2, mean: 15.2}
}
