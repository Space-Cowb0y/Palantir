package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Space-Cowb0y/Palantir/pkg/ui"
)

type model struct {
	choices []string
	cursor  int
}

func initialModel() model {
	return model{
		choices: []string{
			"Start Server (Web + Monitor)",
			"Open Web UI",
			"Run Desktop GUI (Fyne)",
			"Show Plugins (stub)",
			"Exit",
		},
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				fmt.Println("[Starting Web stack: :8080, metrics :9090]")
				go ui.StartWebStack()
			case 1:
				fmt.Println("Open: http://localhost:8080/ui")
			case 2:
				fmt.Println("[Launching Fyne Desktop GUI]")
				go ui.RunFyne()
			case 3:
				fmt.Println("[Plugins list TBD]")
			case 4:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	s := "Choose an option:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = style.Render(">")
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\n(press q to quit)"
	return s
}

func RunCLI() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
