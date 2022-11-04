package task

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/segmentio/ksuid"
)

type markdown struct {
	docType  DocType
	uniqueID ksuid.KSUID
	// localPath string
	filterValue string
	Note        string    `json:"note"`
	Body        string    `json:"body,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func markdownFromText(text string) *markdown {
	md := markdown{
		docType:   InMemoryDoc,
		Body:      text,
		CreatedAt: time.Now(),
	}

	md.generateIDS()
	return &md
}

func (m *markdown) generateIDS() {
	m.uniqueID = ksuid.New()
}

func (m *model) openMarkdown(md *markdown) tea.Cmd {
	var cmd tea.Cmd
	log.Println("opening markdown renderer")
	switch md.docType {
	case InMemoryDoc:
		log.Println("opening in-memory markdown renderer")
		cmd = func() tea.Msg {
			return fetchMarkdownMessage(md)
		}
	default:
		log.Println("invalid markdown type")
	}

	return cmd
}

type fetchMarkdownMessage *markdown
