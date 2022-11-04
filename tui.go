package task

import (
	"context"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-task/task/v3/taskfile"
	"github.com/muesli/termenv"
)

type item struct {
	name    string
	desc    string
	summary string
}

type applicationContext int
type statusMessageTimeoutMsg applicationContext

const (
	pagerContext applicationContext = iota
)

type state int

const (
	stateShowStash state = iota
	stateShowDocument
)

// Title is mandatory needed by DefaultDelegate
func (i item) Title() string {
	return i.name
}

// Description is mandatory needed by DefaultDelegate
func (i item) Description() string {
	return i.desc
}

func (i item) GetName() string {
	return i.name
}

func (i item) FilterValue() string {
	return i.name
}

func (i item) GetSummary() string {
	return i.summary
}

type model struct {
	list     list.Model
	channel  chan string
	state    state
	pager    pagerModel
	fatalErr error
	common   *commonModel
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch key := msg.String(); key {
		case "q", "Q", "ctrl+c":
			close(m.channel)
			return m, tea.Quit
		case "r":
			i := m.list.SelectedItem()
			m.channel <- i.FilterValue()
			close(m.channel)
			return m, tea.Quit
		case " ":
			log.Println("opening summary renderer")
			i := m.list.SelectedItem().(item)
			md := markdownFromText(i.GetSummary())
			md.filterValue = i.GetName()
			cmds = append(cmds, m.openMarkdown(md))

		}
	case tea.WindowSizeMsg:
		// h, v := docStyle.GetFrameSize()
		// m.list.SetSize(msg.Width-h, msg.Height-v)
		log.Printf("new screen size: %d and %d", msg.Width, msg.Height)
		m.common.width = msg.Width
		m.common.height = msg.Height
		topGap, rightGap, bottomGap, leftGap := appStyle.GetPadding()
		m.list.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap)
		m.pager.setSize(msg.Width, msg.Height)
	case fetchMarkdownMessage:
		//
		m.pager.currentDocument = *msg
		cmds = append(cmds, renderWithGlamour(m.pager, msg.Body))
	case contentRenderedMsg:
		log.Println("content rendered")
		m.state = stateShowDocument
	}

	// process state
	switch m.state {
	case stateShowStash:
		newLstModel, lstCmd := m.list.Update(msg)
		m.list = newLstModel
		cmds = append(cmds, lstCmd)
	case stateShowDocument:
		// log.Println("stateShowDocument")
		newPagerModel, cmd := m.pager.Update(msg)
		m.pager = newPagerModel
		cmds = append(cmds, cmd)
	}

	// var c tea.Cmd
	// m.list, c = m.list.Update(msg)
	// cmds = append(cmds, c)
	return m, tea.Batch(cmds...)
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var appStyle = lipgloss.NewStyle().Padding(1, 2)

func (m model) View() string {

	if m.fatalErr != nil {
		return errorView(m.fatalErr, true)
	}

	switch m.state {
	case stateShowDocument:
		return m.pager.View()
	default:
		// return appStyle.Render(m.list.View())
	}

	return docStyle.Render(m.list.View())
}

type Config struct {
	ShowAllFiles    bool
	CharmHost       string `env:"CHARM_HOST" default:"api.charm.sh"`
	Gopath          string `env:"GOPATH"`
	HomeDir         string `env:"HOME"`
	GlamourMaxWidth uint
	GlamourStyle    string
	EnableMouse     bool

	// Which directory should we start from?
	WorkingDirectory string

	// Which document types shall we show?
	DocumentTypes DocTypeSet

	// For debugging the UI
	Logfile              string `env:"GLOW_LOGFILE"`
	HighPerformancePager bool   `env:"GLOW_HIGH_PERFORMANCE_PAGER" default:"true"`
	GlamourEnabled       bool   `env:"GLOW_ENABLE_GLAMOUR" default:"true"`
}

type commonModel struct {
	cfg    Config
	width  int
	height int
}

func (e *Executor) StartTUI(ctx context.Context) {
	// only display tasks with a description
	items := make([]list.Item, 0, len(e.tasksWithDesc()))
	for _, task := range e.tasksWithDesc() {
		items = append(items, item{
			name:    task.Task,
			desc:    task.Desc,
			summary: task.Summary,
		})
	}

	// channel to receive commands from the list
	channel := make(chan string, 1)

	var cfg Config
	cfg.GlamourEnabled = true
	cfg.HighPerformancePager = true
	cfg.WorkingDirectory = e.Dir
	cfg.DocumentTypes = NewDocTypeSet()
	cfg.ShowAllFiles = true
	// cfg.GlamourMaxWidth = width
	cfg.GlamourStyle = "auto"
	cfg.EnableMouse = false
	cfg.DocumentTypes.Add(LocalDoc, InMemoryDoc)
	if cfg.GlamourStyle == "auto" {
		if termenv.HasDarkBackground() {
			cfg.GlamourStyle = "dark"
		} else {
			cfg.GlamourStyle = "light"
		}
	}
	baseModel := commonModel{cfg: cfg}
	m := model{
		list:    list.New(items, list.NewDefaultDelegate(), 0, 0),
		channel: channel,
		pager:   newPagerModel(&baseModel),
		common:  &baseModel,
	}
	m.list.Title = "Available tasks for this project:"
	p := tea.NewProgram(m, tea.WithAltScreen())
	err := p.Start()
	if err != nil {
		fmt.Println("Error running program:", err)
		return
	}

	// wait for a command from the list
	call := <-channel
	if call != "" {
		err = e.RunTask(ctx, taskfile.Call{Task: call})
		if err != nil {
			fmt.Println("Error running task:", err)
			return
		}
	}

}
