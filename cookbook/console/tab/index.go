package tab

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"

// 	"github.com/thuongtruong109/gouse"
// )

// type model struct {
// 	Tabs       []string
// 	TabContent []string
// 	activeTab  int
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch keypress := msg.String(); keypress {
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		case "right", "l", "n", "tab":
// 			m.activeTab = gouse.Min(m.activeTab+1, len(m.Tabs)-1)
// 			return m, nil
// 		case "left", "h", "p", "shift+tab":
// 			m.activeTab = gouse.Max(m.activeTab-1, 0)
// 			return m, nil
// 		}
// 	}

// 	return m, nil
// }

// func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
// 	border := lipgloss.RoundedBorder()
// 	border.BottomLeft = left
// 	border.Bottom = middle
// 	border.BottomRight = right
// 	return border
// }

// var borderColor = "#874BFD"

// var (
// 	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
// 	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
// 	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
// 	highlightColor    = lipgloss.AdaptiveColor{Light: borderColor, Dark: borderColor}
// 	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
// 	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
// 	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
// )

// func (m model) View() string {
// 	doc := strings.Builder{}

// 	var renderedTabs []string

// 	for i, t := range m.Tabs {
// 		var style lipgloss.Style
// 		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
// 		if isActive {
// 			style = activeTabStyle.Copy()
// 		} else {
// 			style = inactiveTabStyle.Copy()
// 		}
// 		border, _, _, _, _ := style.GetBorder()
// 		if isFirst && isActive {
// 			border.BottomLeft = "│"
// 		} else if isFirst && !isActive {
// 			border.BottomLeft = "├"
// 		} else if isLast && isActive {
// 			border.BottomRight = "│"
// 		} else if isLast && !isActive {
// 			border.BottomRight = "┤"
// 		}
// 		style = style.Border(border)
// 		renderedTabs = append(renderedTabs, style.Render(t))
// 	}

// 	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
// 	doc.WriteString(row)
// 	doc.WriteString("\n")
// 	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(m.TabContent[m.activeTab]))
// 	return docStyle.Render(doc.String())
// }

// func Run(tabs []string, tabContent []string) {
// 	m := model{Tabs: tabs, TabContent: tabContent}
// 	if _, err := tea.NewProgram(m).Run(); err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// }
