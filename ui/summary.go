package ui

import tea "github.com/charmbracelet/bubbletea"

type (
	fetchedMarkdownSmrMsg *Summary
)

// Summary
type Summary struct {
	Body string
}

// Command for opening a markdown document in the pager. Note that this also
// alters the model.
func (m *stashModel) openSummaryMarkdown(md *Summary) tea.Cmd {
	var cmd tea.Cmd
	m.viewState = stashStateLoadingDocument

	cmd = loadSummaryMarkdown(md)

	return tea.Batch(cmd, m.spinner.Tick)
}

func loadSummaryMarkdown(md *Summary) tea.Cmd {
	return func() tea.Msg {
		md.Body = string(md.Body)
		return fetchedMarkdownSmrMsg(md)
	}
}
