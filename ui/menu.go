package ui

import (
	"context"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/taskfile"
)

type Item struct {
	Name    string
	Desc    string
	Summary string
}

func (i Item) FilterValue() string {
	return i.Name
}

func (i Item) Title() string {
	return i.Name
}

func (i Item) Description() string {
	return i.Desc
}

func (i Item) GetSummary() string {
	return i.Summary
}

func (i Item) GetName() string {
	return i.Name
}

func Menu(e *task.Executor, ctx context.Context) {
	items := make([]list.Item, 0, len(e.TasksWithDesc()))

	for _, task := range e.TasksWithDesc() {
		items = append(items, Item{
			Name:    task.Name(),
			Desc:    task.Desc,
			Summary: task.Summary,
		})
	}

	var cfg Config
	cfg.GlamourEnabled = true
	cfg.HighPerformancePager = true
	cfg.GlamourStyle = "auto"
	cfg.DocumentTypes = NewDocTypeSet()
	cfg.DocumentTypes.Add(LocalDoc)

	cfg.Logfile = "debug.log"
	_, err := tea.LogToFile(cfg.Logfile, "glow")
	if err != nil {
		log.Println(err)
	}

	log.Println("starting: ")

	// channel to receive task
	channel := make(chan string, 1)

	p := NewProgram(cfg, channel, items)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running tea program:", err)
		return
	}

	log.Println("quit. done.")

	// wait for a task
	call := <-channel
	if call != "" {
		if err := e.RunTask(ctx, taskfile.Call{Task: call}); err != nil {
			fmt.Println("error running task:", err)
			return
		}

	}

}
