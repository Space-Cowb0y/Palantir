package ui
import(
  "fmt"; "net/http"; "time"
  table "github.com/Evertras/bubble-table/table"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)
type EyeRow struct{ Name,Addr,Status,PID,Uptime string }
func QueryEyesOverHTTP(addr string) []EyeRow { _,_ = http.Get("http://"+addr+"/api/eyes"); return []EyeRow{{"greeter","127.0.0.1:50060","running","1234","5s"}} }
type model struct{ table table.Model }
func NewManagerModel(fetch func()([]EyeRow,error)) model{ cols:=[]table.Column{{Title:"Name",Width:20},{Title:"Addr",Width:20},{Title:"Status",Width:10},{Title:"PID",Width:7},{Title:"Uptime",Width:10}}; m:=model{ table: table.New(cols).WithRows([]table.Row{}).Border(lipgloss.NormalBorder()).WithPageSize(10)}; return m }
func (m model) Init() tea.Cmd { return tea.Tick(time.Second, func(time.Time) tea.Msg { return "tick" }) }
func (m model) Update(msg tea.Msg)(tea.Model,tea.Cmd){ switch msg.(type){ case tea.KeyMsg: return m, tea.Quit; case string: return m, tea.Tick(time.Second, func(time.Time) tea.Msg { return "tick" }) }; return m, nil }
func (m model) View() string { return lipgloss.NewStyle().Padding(1).Render("Sentinel — Eyes manager\n\n"+m.table.View()+"\n\nPress any key to exit.") }
func Run(m model) error { p:=tea.NewProgram(m); if _,err:=p.Run(); err!=nil{ return fmt.Errorf("tui: %w",err)}; return nil }