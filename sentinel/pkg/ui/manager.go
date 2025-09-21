package ui

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lgtable "github.com/charmbracelet/lipgloss/table"
)

type EyeRow struct {
	Name, Addr, Status, PID, Uptime string
}

var eyesState = map[string]time.Time{} // chave = PID ou Addr

func QueryEyesOverHTTP(addr string) []EyeRow {
	resp, err := http.Get("http://" + addr + "/api/eyes")
	if err != nil {
		// em produção, faça tratamento de erro
		return nil
	}
	defer resp.Body.Close()

	// TODO: parse JSON real, por enquanto vamos simular
	pid := "1234"
	name := "greeter"

	// se nunca vimos esse PID antes, salva StartTime
	if _, ok := eyesState[pid]; !ok {
		eyesState[pid] = time.Now()
	}

	uptime := time.Since(eyesState[pid]).Truncate(time.Second)

	return []EyeRow{{
		Name:   name,
		Addr:   addr,
		Status: "running",
		PID:    pid,
		Uptime: uptime.String(),
	}}
}

type model struct {
	table *lgtable.Table
	fetch func() ([]EyeRow, error)
	err   error
}

// ---- helpers ----

func makeTable() *lgtable.Table {
	// Estilos
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true)

	cellStyle := lipgloss.NewStyle().Padding(0, 1)

	// Constrói a tabela base (sem linhas ainda)
	t := lgtable.New().
		Border(lipgloss.RoundedBorder()).
		Headers("Name", "Addr", "Status", "PID", "Uptime").
		StyleFunc(func(row, col int) lipgloss.Style {
			// primeira linha é o header (row==0) — lipgloss/table já aplica header separadamente,
			// mas podemos padronizar as células e reforçar o estilo do cabeçalho:
			if row == 0 {
				return headerStyle
			}
			return cellStyle
		})

	// Opcional: largura total (o lib autoajusta colunas)
	// t = t.Width(72)

	return t
}

func toRowStrings(rr []EyeRow) [][]string {
	out := make([][]string, 0, len(rr))
	for _, r := range rr {
		out = append(out, []string{r.Name, r.Addr, r.Status, r.PID, r.Uptime})
	}
	return out
}

// ---- model ----

func NewManagerModel(fetch func() ([]EyeRow, error)) model {
	t := makeTable()
	return model{
		table: t,
		fetch: fetch,
	}
}

func (m model) Init() tea.Cmd {
	// Atualiza imediatamente e depois de 1 em 1 segundo
	return tea.Batch(
		func() tea.Msg {
			return tickMsg(time.Now())
		},
	)
}

type tickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// Qualquer tecla sai (igual ao seu comportamento)
		return m, tea.Quit

	case tea.WindowSizeMsg:
		// Se quiser ajustar largura conforme terminal:
		// m.table = m.table.Width(msg.Width - 4)
		return m, nil

	case tickMsg:
		// a cada tick, tenta buscar e atualizar linhas
		if m.fetch != nil {
			rows, err := m.fetch()
			if err != nil {
				m.err = err
			} else {
				m.err = nil
				m.table = m.table.Rows(toRowStrings(rows)...)
			}
		}
		// agenda próximo tick
		return m, tea.Tick(30*time.Second, func(time.Time) tea.Msg { return tickMsg(time.Now()) })

	default:
		return m, nil
	}
}

func (m model) View() string {
	title := lipgloss.NewStyle().Bold(true).Render("Sentinel — Eyes manager")
	body := m.table.Render()
	footer := "Press any key to exit."

	if m.err != nil {
		errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
		footer = errStyle.Render("error: "+m.err.Error()) + "\n" + footer
	}

	return lipgloss.NewStyle().Padding(1).Render(
		title + "\n\n" + body + "\n\n" + footer,
	)
}

func Run(m model) error {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("tui: %w", err)
	}
	return nil
}
