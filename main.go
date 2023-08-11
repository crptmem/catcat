package main
import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  _ "github.com/spf13/viper"
  "github.com/charmbracelet/bubbles/spinner"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
  spinner  spinner.Model
}
var m model

func (m model) Init() tea.Cmd {
  s := spinner.New()
	s.Spinner = spinner.Dot
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    if msg.String() == "ctrl+c" {
      return m, tea.Quit
    }
    if msg.String() == "enter" {
      var Item = m.list.SelectedItem().(item)
      var Entry = getGameEntry(Item.title)
      var err = launchGame(Entry["executablepath"].(string), Entry["wrappercommand"].(string), Entry["discordpresence"].(bool))
      if err != nil {
        m.list.NewStatusMessage("Executable does not exist")
      }
      return m, nil
    }
  case tea.WindowSizeMsg:
    h, v := docStyle.GetFrameSize()
    m.list.SetSize(msg.Width-h, msg.Height-v)
  }

  var cmd tea.Cmd
  m.list, cmd = m.list.Update(msg)
  return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
  readConfig() 
  var GameEntries = getGameEntries()
  
  items := []list.Item{}

  for _, entry := range GameEntries {
    entryMap := entry.(map[string]interface{})
    if executablePath, ok := entryMap["executablepath"].(string); ok {
      if name, ok := entryMap["name"].(string); ok {
        fmt.Printf("\nName: %s, ExecutablePath: %s", name, executablePath)
        items = append(items, item{title: name, desc: executablePath + " | " + entryMap["wrappercommand"].(string)})
      } else {
        fmt.Println("Name is nil or not a string")
      }
    } else {
      fmt.Println("ExecutablePath is nil or not a string")
    }
  }

  m = model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
  m.list.Title = "catcat v0.1"
  m.list.SetShowStatusBar(true)
  p := tea.NewProgram(m, tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
}

