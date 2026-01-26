package tui

import (
	"ATAD_PFCLIManager/internal/core/transaction"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var styleHeader = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00ADD8"))

type model struct {
	transactions []transaction.Transaction
	cursor       int
	err          error
}

func NewModel(txs []transaction.Transaction) model {
	return model{transactions: txs}
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
			if m.cursor < len(m.transactions)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := styleHeader.Render("--- BROWSE TRANSACTIONS (q to quit) ---\n\n")
	for i, tx := range m.transactions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s [%s] %-15s %10.2f RON\n", cursor, tx.Date.Format("02 Jan"), tx.Description, tx.Amount)
	}
	return s
}

func StartTUI(txs []transaction.Transaction) error {
	p := tea.NewProgram(NewModel(txs))
	_, err := p.Run()
	return err
}
