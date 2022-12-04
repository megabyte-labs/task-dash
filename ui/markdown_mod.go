package ui

import tea "github.com/charmbracelet/bubbletea"

type (
	fetchedMarkdownSmrMsg *Markdown
)

// Markdown is wrap of markdown
type Markdown struct {
	markdown
	Body string
}

// Command for opening a markdown document in the pager. Note that this also
// alters the model.
func (m *stashModel) openSummaryMarkdown(md *Markdown) tea.Cmd {
	var cmd tea.Cmd
	m.viewState = stashStateLoadingDocument

	cmd = loadSummaryMarkdown(md)

	return tea.Batch(cmd, m.spinner.Tick)
}

func loadSummaryMarkdown(md *Markdown) tea.Cmd {
	return func() tea.Msg {
		md.Body = string(md.Body)
		return fetchedMarkdownSmrMsg(md)
	}
}
