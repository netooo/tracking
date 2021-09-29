package tracking

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	taskFile = "task.json"
	taskPath = filepath.Join(cacheDir, taskFile)
	tasks    []*Task
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContentLine int    `json:"content_line"`
}

func (t *Task) Add() error {
	taskList, err := TaskRead()
	if err != nil {
		return err
	}

	for _, task := range taskList {
		if t.ID == task.ID {
			return errors.New("the task already exists")
		}
	}

	newTasks := append(taskList, t)
	buf, err := json.Marshal(newTasks)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(taskPath, os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}
