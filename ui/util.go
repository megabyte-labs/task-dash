package ui

import (
	"os"
	"path"

	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/taskfile"
	"gopkg.in/yaml.v3"
)

// Taskfile represents a Taskfile.yml
type Taskfile struct {
	Bookmarked taskfile.Tasks
}

// UnmarshalYAML implements yaml.Unmarshaler interface
// func (tf *Taskfile) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	var taskfile struct {
// 		Bookmarked taskfile.Tasks
// 	}
// 	if err := unmarshal(&taskfile); err != nil {
// 		return err
// 	}
// 	tf.Bookmarked = taskfile.Bookmarked
// 	return nil
// }

func readTaskfile(file string) (*Taskfile, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var t Taskfile
	return &t, yaml.NewDecoder(f).Decode(&t)
}

func parseBookmark(e *task.Executor) (*Taskfile, error) {
	t, err := readTaskfile(path.Join(e.Dir, e.Entrypoint))
	if err != nil {
		return nil, err
	}

	for name, task := range t.Bookmarked {
		task.Task = name
	}

	return t, nil
}
